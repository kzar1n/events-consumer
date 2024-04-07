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

func (r *PaymentRepository) FindById(Id string) (*entity.Payment, error) {
	payment := &entity.Payment{}
	sqlStatement := `SELECT id_payment, id_account, payment_type, payment_date, payment_value FROM payments WHERE id_payment=?`
	return payment, r.DB.QueryRow(sqlStatement, Id).Scan(&payment.Id, &payment.IdAccount, &payment.PaymentType, &payment.PaymentDate, &payment.PaymentValue)
}

func (r *PaymentRepository) Insert(payment *entity.Payment) error {
	sqlStatement := `INSERT INTO payments (id_payment, id_account, payment_type, payment_date, payment_value) VALUES (?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(sqlStatement, payment.Id, payment.IdAccount, payment.PaymentType, payment.PaymentDate, payment.PaymentValue)
	return err
}

func (r *PaymentRepository) UpdateById(Id string, newPayment *entity.Payment) error {
	sqlStatement := `UPDATE payments SET id_account = ?, payment_type = ?, payment_date = ?, payment_value = ?, updated_at = current_timestamp, update_count = update_count + 1 WHERE id_payment = ?`
	_, err := r.DB.Exec(sqlStatement, newPayment.IdAccount, newPayment.PaymentType, newPayment.PaymentDate, newPayment.PaymentValue, Id)
	return err
}
