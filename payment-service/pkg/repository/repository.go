package repository

import (
	"ms-practice/payment-service/pkg/repository/payment"

	"gorm.io/gorm"
)

type Repository struct {
	PaymentRepo payment.PaymentRepoInterface
}

func NewRepository(db *gorm.DB) *Repository {
	paymentRepo := payment.NewPaymentRepo(db)
	return &Repository{
		PaymentRepo: paymentRepo,
	}
}
