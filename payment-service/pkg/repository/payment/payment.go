package payment

import (
	"context"
	"ms-practice/payment-service/pkg/model"

	"gorm.io/gorm"
)

type PaymentRepoInterface interface {
	GetPayment(ctx context.Context, userID int32, paymentID int32) (*model.Payment, error)
}

type paymentRepo struct {
	db *gorm.DB
}

func NewPaymentRepo(db *gorm.DB) PaymentRepoInterface {
	return &paymentRepo{
		db: db,
	}
}

func (r *paymentRepo) GetPayment(ctx context.Context, userID int32, paymentID int32) (*model.Payment, error) {
	var booking model.Payment
	err := r.db.WithContext(ctx).Preload("PaymentHistories").First(&booking).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}
