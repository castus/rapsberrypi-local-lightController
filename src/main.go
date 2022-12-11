package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"raspberrypi.local/lightController/mqttHandler"
)

const (
	TOPIC     = "leds"
	QOS       = 1
	CLIENT_ID = "triggerApi"
)

var SERVER_ADDRESS = os.Getenv("MQTT_SERVER_ADDRESS")

var h = mqttHandler.NewHandler()

func main() {
	mqtt.ERROR = log.New(os.Stdout, "[ERROR] ", 0)
	mqtt.CRITICAL = log.New(os.Stdout, "[CRITICAL] ", 0)
	mqtt.WARN = log.New(os.Stdout, "[WARN]  ", 0)
	mqtt.DEBUG = log.New(os.Stdout, "[DEBUG] ", 0)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(SERVER_ADDRESS)
	opts.SetClientID(CLIENT_ID)
	opts.SetOrderMatters(false)       // Allow out of order messages (use this option unless in order delivery is essential)
	opts.ConnectTimeout = time.Second // Minimal delays on connect
	opts.WriteTimeout = time.Second   // Minimal delays on writes
	opts.KeepAlive = 10               // Keepalive every 10 seconds so we quickly detect network outages
	opts.PingTimeout = time.Second    // local broker so response should be quick
	opts.ConnectRetry = true
	opts.AutoReconnect = true
	opts.DefaultPublishHandler = func(_ mqtt.Client, msg mqtt.Message) {
		fmt.Printf("UNEXPECTED MESSAGE: %s\n", msg)
	}
	opts.OnConnectionLost = func(cl mqtt.Client, err error) {
		fmt.Println("connection lost")
	}
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("connection established")
		t := c.Subscribe(TOPIC, QOS, handle)
		go func() {
			_ = t.Wait() // Can also use '<-t.Done()' in releases > 1.2.0
			if t.Error() != nil {
				fmt.Printf("ERROR SUBSCRIBING: %s\n", t.Error())
			} else {
				fmt.Println("subscribed to: ", TOPIC)
			}
		}()
	}
	opts.OnReconnecting = func(mqtt.Client, *mqtt.ClientOptions) {
		fmt.Println("attempting to reconnect")
	}

	client := mqtt.NewClient(opts)
	client.AddRoute(TOPIC, handle)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	fmt.Println("Connection is up")

	// Messages will be delivered asynchronously, so we just need to wait for a signal to shut down
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	<-sig
	fmt.Println("signal caught - exiting")
	client.Disconnect(1000)
	fmt.Println("shutdown complete")
}

func handle(_ mqtt.Client, msg mqtt.Message) {
	var m mqttHandler.Message
	err := json.Unmarshal(msg.Payload(), &m)
	// Don't react on every message
	if err != nil {
		return
	}

	fmt.Printf("received message: %s\n", msg.Payload())
	h.Handle(m)
}
