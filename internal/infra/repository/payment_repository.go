package repository

import (
	"database/sql"

	"github.com/kzar1n/events-consumer/internal/entity"
)

type PaymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{
		DB: db,
	}
}

func (r *PaymentRepository) Save(payment *entity.Payment) error {
	// Save payment to database
	return nil
}
