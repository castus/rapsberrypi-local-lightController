#!/bin/bash

go mod init raspberrypi.local/lightController
go get github.com/eclipse/paho.mqtt.golang
go get github.com/gorilla/websocket
go get golang.org/x/net/proxy

go run raspberrypi.local/lightController
