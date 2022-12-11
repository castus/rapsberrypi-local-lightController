package mqttHandler

import (
	"raspberrypi.local/lightController/alexaTrigger"
	"raspberrypi.local/lightController/lightController"
)

type Handler struct {
	alexaTrigger *alexaTrigger.AlexaTrigger
}

func NewHandler() *Handler {
	trigger := alexaTrigger.NewAlexaTrigger()
	return &Handler{
		alexaTrigger: trigger,
	}
}

func (h *Handler) Handle(m Message) {
	if m.IsLightOn {
		triggerKey := lightController.GetTriggerKey()
		eventChan := make(chan string)
		go h.alexaTrigger.DebounceTrigger(eventChan)
		eventChan <- triggerKey
	}
}

type Message struct {
	IsLightOn bool
	Place     string
}
