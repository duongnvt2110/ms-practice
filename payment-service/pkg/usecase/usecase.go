package usecase

import (
	"ms-practice/payment-service/pkg/repository"
	"ms-practice/payment-service/pkg/usecase/payment"
)

type Usecase struct {
	PaymentUC payment.PaymentUsecaseInterface
}

func NewUsecase(repo *repository.Repository) *Usecase {
	paymentUC := payment.NewPaymentUsecase(repo.PaymentRepo)
	return &Usecase{
		PaymentUC: paymentUC,
	}
}
