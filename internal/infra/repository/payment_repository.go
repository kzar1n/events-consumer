package repository

import (
	"database/sql"
	"fmt"

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

func (r *PaymentRepository) FindByID(id string) (*entity.Payment, error) {
	fmt.Printf("Finding payment by ID: %s\n", id)
	payment := &entity.Payment{}
	query := `SELETEC id_payment, id_account, payment_type, payment_date, payment_value FROM payments WHERE id = $1`
	err := r.DB.QueryRow(query, id).Scan(&payment.ID, &payment.IdAccount, &payment.PaymentType, &payment.PaymentDate, &payment.PaymentValue)

	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (r *PaymentRepository) Update(payment *entity.Payment, newPayment *entity.Payment) error {
	// Update payment in database
	return nil
}
