package payment

import (
	"context"
	"fmt"
	"ms-practice/payment-service/pkg/model"
	"ms-practice/payment-service/pkg/repository/payment"
	"ms-practice/pkg/event"
	"time"

	"github.com/google/uuid"
)

type PaymentUsecaseInterface interface {
	GetPayment(ctx context.Context, userID int32, paymentID int32) (*model.Payment, error)
	ProcessPayment(ctx context.Context, payload event.BookingPayload) (*model.Payment, error)
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

func (u *paymentUsecase) ProcessPayment(ctx context.Context, payload event.BookingPayload) (*model.Payment, error) {
	now := time.Now()
	payment := &model.Payment{
		IdempotencyKey: fmt.Sprintf("booking-%d", payload.OrderID),
		UserId:         payload.UserID,
		BookingID:      payload.OrderID,
		TransactionID:  uuid.NewString(),
		PaymentCode:    fmt.Sprintf("PMT-%s", uuid.NewString()),
		Amount:         int(payload.Amount),
		Status:         "succeeded",
		Provider:       "mock",
		PaidAt:         &now,
	}

	if err := u.paymentRepo.CreatePayment(ctx, payment); err != nil {
		return nil, err
	}

	return payment, nil
}
