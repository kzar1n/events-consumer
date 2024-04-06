package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/kzar1n/events-consumer/internal/entity"
)

type EventDTO struct {
	ID        string    `json:"id"`
	IDPayment string    `json:"id_payment"`
	EventType string    `json:"event_type"`
	EventDate time.Time `json:"event_date"`
}

func NewEventDTO(event entity.Event) *EventDTO {
	return &EventDTO{
		ID:        event.ID,
		IDPayment: event.IDPayment,
		EventType: event.EventType,
		EventDate: event.EventDate,
	}
}

func EventProcessor(eventDto *EventDTO) {
	fmt.Printf("Processing event: %v\n", eventDto)
	payment, err := FindPaymentByID(eventDto.IDPayment)

	// TODO improve this process with some interface
	if err == sql.ErrNoRows {
		SavePayment(FindLegacyPaymentByID(eventDto.IDPayment))
	} else if err != nil {
		panic(err)
	} else {
		UpdatePayment(payment, FindLegacyPaymentByID(eventDto.IDPayment))
	}
}
