package payment

import (
	"context"
	"ms-practice/payment-service/pkg/model"
	"ms-practice/payment-service/pkg/repository/payment"
)

type PaymentUsecaseInterface interface {
	GetPayment(ctx context.Context, userID int32, paymentID int32) (*model.Payment, error)
}

type paymentUsecase struct {
	paymentRepo payment.PaymentRepoInterface
}

func NewPaymentUsecase(paymentRepo payment.PaymentRepoInterface) PaymentUsecaseInterface {
	return &paymentUsecase{
		paymentRepo: paymentRepo,
	}
}

func (u *paymentUsecase) GetPayment(ctx context.Context, userID int32, paymentID int32) (*model.Payment, error) {
	return u.paymentRepo.GetPayment(ctx, userID, paymentID)
}
