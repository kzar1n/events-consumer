package repository

import (
	"database/sql"

	"github.com/kzar1n/events-consumer/internal/entity"
)

type LegacyPaymentRepository struct {
	DB *sql.DB
}

func NewLegacyPaymentRepository(db *sql.DB) *LegacyPaymentRepository {
	return &LegacyPaymentRepository{
		DB: db,
	}
}

func (t *LegacyPaymentRepository) FindLegacyPaymentById(Id string) (*entity.LegacyPayment, error) {
	legacyPayment := &entity.LegacyPayment{}
	sqlStatement := `SELECT id_payment, id_account, payment_type, payment_date, payment_value FROM payments WHERE id_payment=$1`
	return legacyPayment, t.DB.QueryRow(sqlStatement, Id).Scan(&legacyPayment.Id, &legacyPayment.IdAccount, &legacyPayment.PaymentType, &legacyPayment.PaymentDate, &legacyPayment.PaymentValue)
}

func (t *LegacyPaymentRepository) Insert(legacyPayment *entity.LegacyPayment) error {
	sqlStatement := `INSERT INTO payments (id_payment, id_account, payment_type, payment_date, payment_value, payment_status) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := t.DB.Exec(sqlStatement, legacyPayment.Id, legacyPayment.IdAccount, legacyPayment.PaymentType, legacyPayment.PaymentDate, legacyPayment.PaymentValue, legacyPayment.PaymentStatus)
	return err
}
