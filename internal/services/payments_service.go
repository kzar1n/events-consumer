package services

import (
	"fmt"

	"github.com/kzar1n/events-consumer/internal/entity"
	"github.com/kzar1n/events-consumer/internal/infra/database"
	"github.com/kzar1n/events-consumer/internal/infra/repository"
)

func FindPaymentByID(id string) (*entity.Payment, error) {
	fmt.Printf("Finding payment by ID: %s\n", id)
	payment, err := repository.NewPaymentRepository(database.DB).FindByID(id)

	if err != nil {
		return nil, err
	}
	return payment, nil
}

func SavePayment(payment *entity.Payment) error {
	return repository.NewPaymentRepository(database.DB).Save(payment)
}

func UpdatePayment(payment *entity.Payment, newPayment *entity.Payment) error {
	return repository.NewPaymentRepository(database.DB).Update(payment, newPayment)
}

func FindLegacyPaymentByID(id string) *entity.Payment {
	return entity.NewPaymentFromLegacy(id)
}
