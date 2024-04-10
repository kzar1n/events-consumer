package entity

import (
	"time"
)

type LegacyPayment struct {
	Id            string
	IdAccount     string
	PaymentType   string
	PaymentDate   time.Time
	PaymentValue  float64
	PaymentStatus string
}

func NewLegacyPayment(id, idAccount, paymentType string) *LegacyPayment {
	return &LegacyPayment{
		Id:            id,
		IdAccount:     idAccount,
		PaymentType:   paymentType,
		PaymentDate:   time.Now(),
		PaymentValue:  100.00,
		PaymentStatus: "PENDING",
	}
}
