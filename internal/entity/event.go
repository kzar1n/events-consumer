package entity

import "time"

type Event struct {
	ID         string
	ID_payment string
	Event_type string
	Event_date time.Time
}

func NewEvent(id string, id_payment string, event_type string, event_date time.Time) *Event {
	return &Event{
		ID:         id,
		ID_payment: id_payment,
		Event_type: event_type,
		Event_date: event_date,
	}
}
