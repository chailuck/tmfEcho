package party

import (
	"time"
	"tmfEcho/internal/log"
)

type IndividualEventHandlerData struct {
	CorrelationId   string          `json:"correlationId,omitempty"`
	Domain          string          `json:"domain,omitempty"`
	EventId         string          `json:"eventId,omitempty"`
	EventTime       string          `json:"eventTime,omitempty"`
	EventType       string          `json:"eventType,omitempty"`
	Priority        string          `json:"priority,omitempty"`
	TimeOcurred     string          `json:"timeOcurred,omitempty"`
	Title           string          `json:"title,omitempty"`
	Event           individual      `json:"event,omitempty"`
	ReportingSystem reportingSystem `json:"reportingSystem,omitempty"`
	Source          source          `json:"source,omitempty"`
	Type            string          `json:"@type,omitempty"`
	BaseType        string          `json:"@baseType,omitempty"`
}

type individual struct {
	Individual IndividualData `json:"individual,omitempty"`
}

type reportingSystem struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Type         string `json:"@type,omitempty"`
	ReferredType string `json:"@referredType,omitempty"`
}

type source struct {
	Id           string `json:"id,omitempty"`
	Name         string `json:"name,omitempty"`
	Type         string `json:"@type,omitempty"`
	ReferredType string `json:"@referredType,omitempty"`
}

type EventHandler struct {
	indyData IndividualEventHandlerData
}

func (e *EventHandler) Initialize(corrId string, eventId string) {
	e.indyData = IndividualEventHandlerData{}
	e.indyData.CorrelationId = corrId
	e.indyData.Domain = "CRM"
	e.indyData.EventId = eventId

}

func (e *EventHandler) IndividualCreateEventTrigger(IndividualData IndividualData, lg log.LogTracing) {
	e.indyData.Type = "IndividualCreateEvent"
	e.indyData.BaseType = "Event"
	e.indyData.EventType = "IndividualCreateEvent"
	e.indyData.Title = "IndividualCreateEvent"
	e.indyData.Event.Individual = IndividualData
	currTime := time.Now()
	currTimeStr := currTime.Format("01-02-2006T15:04:05.999Z")
	e.indyData.EventTime = currTimeStr
	e.indyData.TimeOcurred = currTimeStr
	e.Handler()
}

func (e *EventHandler) Handler() {
	// do something

}
