package mqttHandler

import (
	"raspberrypi.local/lightController/alexaTrigger"
	"raspberrypi.local/lightController/lightController"
)

type Handler struct {
	alexaTrigger *alexaTrigger.AlexaTrigger
}

type Message struct {
	IsLightOn bool
	Place     string
}

func NewHandler() *Handler {
	trigger := alexaTrigger.NewAlexaTrigger()
	return &Handler{
		alexaTrigger: trigger,
	}
}

func (h *Handler) Handle(m Message) {
	h.sendOnOffTrigger(m)
	h.sendTrigger(lightController.GetTriggerKeyMatchingTimeOfADay())
}

func (h *Handler) sendTrigger(triggerKey string) {
	eventChan := make(chan string)
	go h.alexaTrigger.DebounceTrigger(eventChan)
	eventChan <- triggerKey
}

func (h *Handler) sendOnOffTrigger(m Message) {
	if m.Place == PlaceTrees.String() || m.Place == PlaceTV.String() {
		if m.IsLightOn {
			h.sendTrigger("trigger-on")
		} else {
			h.sendTrigger("trigger-off")
		}
	}
}
