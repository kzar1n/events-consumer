package entity

import "time"

type PaymentRepository interface {
	Save(payment *Payment) error
}

type Payment struct {
	ID           string
	IdAccount    string
	PaymentType  string
	PaymentDate  time.Time
	PaymentValue float64
}

func NewPayment(id, idAccount, paymentType string, paymentDate time.Time, paymentValue float64) *Payment {
	return &Payment{
		ID:           id,
		IdAccount:    idAccount,
		PaymentType:  paymentType,
		PaymentDate:  paymentDate,
		PaymentValue: paymentValue,
	}
}
