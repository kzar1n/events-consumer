package entity

import "time"

type Event struct {
	Id        string
	IdPayment string
	EventType string
	EventDate time.Time
}

func NewEvent(Id string, IdPayment string, EventType string, EventDate time.Time) *Event {
	return &Event{
		Id:        Id,
		IdPayment: IdPayment,
		EventType: EventType,
		EventDate: EventDate,
	}
}
