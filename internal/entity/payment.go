package entity

import "time"

type PaymentRepository interface {
	Save(payment *Payment) error
}

type Payment struct {
	Id           string
	IdAccount    string
	PaymentType  string
	PaymentDate  time.Time
	PaymentValue float64
}

func NewPayment(Id, IdAccount, paymentType string, paymentDate time.Time, paymentValue float64) *Payment {
	return &Payment{
		Id:           Id,
		IdAccount:    IdAccount,
		PaymentType:  paymentType,
		PaymentDate:  paymentDate,
		PaymentValue: paymentValue,
	}
}
