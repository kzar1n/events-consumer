package entity

import "time"

func NewPaymentFromLegacy(id string) *Payment {
	return &Payment{
		ID:           id,
		IdAccount:    "1",
		PaymentType:  "credit",
		PaymentDate:  time.Now(),
		PaymentValue: 100.00,
	}
}
