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
	sqlStatement := `SELECT id_payment, id_account, payment_type, payment_date, payment_value FROM payments WHERE id_payment=?`
	return legacyPayment, t.DB.QueryRow(sqlStatement, Id).Scan(&legacyPayment.Id, &legacyPayment.IdAccount, &legacyPayment.PaymentType, &legacyPayment.PaymentDate, &legacyPayment.PaymentValue)
}
