package events

import (
	"strconv"
)

type listener struct {
	Type           string `json:"@type"`
	BaseType       string `json:"@baseType,omitempty"`
	SchemaLocation string `json:"@schemaLocation,omitempty"`
	Callback       string `json:"callback,omitempty"`
	Query          string `json:"query,omitempty"`
	Id             string `json:"id,omitempty"`
}

type EventListener struct {
	notifiers map[string]listener

	maxID int64
}

func Initialize() *EventListener {
	return &EventListener{make(map[string]listener), 1}
}

func (u *EventListener) RegisterNotifier(request listener) {
	u.notifiers[strconv.FormatInt(u.maxID, 10)] = request
	u.maxID++
}

func (u *EventListener) RemoveNotifier(id string) {
	delete(u.notifiers, id)
}

/*
func (u *EventListener) Notify(payload interface{}) {
	for _, n := range u.listeners {


	}
}
*/
