package alexaTrigger

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type AlexaTrigger struct {
}

func NewAlexaTrigger() *AlexaTrigger {
	return &AlexaTrigger{}
}

func (m *AlexaTrigger) DebounceTrigger(input chan string) {
	var item string
	interval := time.Millisecond * 1500
	timer := time.NewTimer(interval)
	for {
		select {
		case item = <-input:
			timer.Reset(interval)
			fmt.Println("Trigger debounced")
		case <-timer.C:
			if item != "" {
				fmt.Println("Trigger run")
				m.Trigger(item)
			}
		}
	}
}

func (m *AlexaTrigger) Trigger(key string) {
	URL := fmt.Sprintf("https://api.voicemonkey.io/trigger?access_token=%s&secret_token=%s&monkey=%s",
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("SECRET_TOKEN"),
		key)

	if os.Getenv("SHOULD_TRIGGER_ALEXA") == "false" {
		fmt.Printf("Alexa trigger not send to URL: %s\n", URL)
		return
	} else {
		fmt.Printf("Alexa trigger send to URL: %s\n", URL)
	}

	resp, err := http.Get(URL)

	if err != nil {
		fmt.Printf("Request Failed: %s", err)
		panic("Monkey trigger request Failed")
	}
	defer resp.Body.Close()

	fmt.Println("Monkey trigger request Success")
	if err != nil {
		fmt.Printf("Reading body failed: %s", err)
		panic("Reading Monkey trigger body failed")
	}
}
