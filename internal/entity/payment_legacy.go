package entity

import (
	"time"

	"github.com/google/uuid"
)

type LegacyPayment struct {
	Id           string
	IdAccount    string
	PaymentType  string
	PaymentDate  time.Time
	PaymentValue float64
}

func NewLegacyPayment(Id string) *LegacyPayment {
	return &LegacyPayment{
		Id:           Id,
		IdAccount:    uuid.New().String(),
		PaymentType:  "legacy",
		PaymentDate:  time.Now(),
		PaymentValue: 100.00,
	}
}
