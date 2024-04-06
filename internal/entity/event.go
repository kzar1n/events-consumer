package entity

import "time"

type Event struct {
	ID        string
	IDPayment string
	EventType string
	EventDate time.Time
}

func NewEvent(id string, IDPayment string, EventType string, EventDate time.Time) *Event {
	return &Event{
		ID:        id,
		IDPayment: IDPayment,
		EventType: EventType,
		EventDate: EventDate,
	}
}
