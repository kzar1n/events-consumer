package services

import (
	"fmt"
	"time"

	"github.com/kzar1n/events-consumer/internal/entity"
)

type EventDTO struct {
	Id        string    `json:"Id"`
	IdPayment string    `json:"Id_payment"`
	EventType string    `json:"event_type"`
	EventDate time.Time `json:"event_date"`
}

func NewEventDTO(event *entity.Event) *EventDTO {
	return &EventDTO{
		Id:        event.Id,
		IdPayment: event.IdPayment,
		EventType: event.EventType,
		EventDate: event.EventDate,
	}
}

func (t *EventDTO) EventProcessor() {
	paymentService := NewPaymentService()
	legacyPayment, err := paymentService.FindLegacyPaymentById(t.IdPayment)
	if err != nil {
		fmt.Println("Legacy Payment not found! -> ", t.IdPayment, " error: ", err)
		panic(err)
	}

	if err := paymentService.SaveOrUpdatePaymentFromLegacy(legacyPayment); err != nil {
		panic(err)
	}
}
