package services

import (
	"database/sql"

	"github.com/kzar1n/events-consumer/internal/entity"
	"github.com/kzar1n/events-consumer/internal/infra/database"
	"github.com/kzar1n/events-consumer/internal/infra/repository"
)

type PaymentService struct {
	PaymentRepository       *repository.PaymentRepository
	LegacyPaymentRepository *repository.LegacyPaymentRepository
}

func NewPaymentService() *PaymentService {
	return &PaymentService{
		PaymentRepository:       repository.NewPaymentRepository(database.MysqlDB),
		LegacyPaymentRepository: repository.NewLegacyPaymentRepository(database.PostgreDB),
	}
}

func (t *PaymentService) FindPaymentById(Id string) (*entity.Payment, error) {
	return t.PaymentRepository.FindById(Id)
}

func (t *PaymentService) SaveOrUpdatePaymentFromLegacy(legacyPayment *entity.LegacyPayment) error {
	payment := entity.NewPayment(legacyPayment.Id, legacyPayment.IdAccount, legacyPayment.PaymentType, legacyPayment.PaymentDate, legacyPayment.PaymentValue)

	switch _, err := t.FindPaymentById(payment.Id); err {
	case sql.ErrNoRows:
		return t.PaymentRepository.Insert(payment)
	case nil:
		return t.PaymentRepository.UpdateById(payment.Id, payment)
	default:
		panic(err)
	}
}

func (t *PaymentService) FindLegacyPaymentById(Id string) (*entity.LegacyPayment, error) {
	return t.LegacyPaymentRepository.FindLegacyPaymentById(Id)
}

func (t *PaymentService) SaveLegacyPayment(legacyPayment *entity.LegacyPayment) error {
	return t.LegacyPaymentRepository.Insert(legacyPayment)
}
