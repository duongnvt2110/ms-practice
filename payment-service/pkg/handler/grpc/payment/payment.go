package payment

import (
	"context"
	"ms-practice/payment-service/pkg/model"
	paymentusecase "ms-practice/payment-service/pkg/usecase/payment"
	"ms-practice/proto/gen"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type PaymentHandler struct {
	gen.UnimplementedPaymentServiceServer
	paymentUC paymentusecase.PaymentUsecaseInterface
}

func NewPaymentGrpcHandler(paymentUC paymentusecase.PaymentUsecaseInterface) *PaymentHandler {
	return &PaymentHandler{paymentUC: paymentUC}
}

func (h *PaymentHandler) GetPayment(ctx context.Context, req *gen.GetPaymentRequest) (*gen.GetPaymentResponse, error) {
	payment, err := h.paymentUC.GetPayment(ctx, req.Id, req.PaymentId)
	if err != nil {
		return nil, err
	}

	return &gen.GetPaymentResponse{
		Payment: &gen.Payment{
			Id:             int32(payment.Id),
			UserId:         int32(payment.UserId),
			BookingId:      int32(payment.BookingID),
			Prices:         float32(payment.Prices),
			CreatedAt:      timestamppb.New(payment.CreatedAt),
			UpdatedAt:      timestamppb.New(payment.UpdatedAt),
			PaymentHistory: toProtoPaymentHistory(payment.PaymentHistories),
		},
	}, nil
}

func toProtoPaymentHistory(histories []model.PaymentHistory) *gen.PaymentHistory {
	if len(histories) == 0 {
		return nil
	}
	history := histories[0]
	var paidAt string
	if history.PaidAt != nil {
		paidAt = history.PaidAt.Format(time.RFC3339)
	}

	return &gen.PaymentHistory{
		Id:        int32(history.Id),
		PaymentId: int32(history.PaymentID),
		Status:    history.Status,
		Logs:      history.Logs,
		PaidAt:    paidAt,
		CreatedAt: timestamppb.New(history.CreatedAt),
		UpdatedAt: timestamppb.New(history.UpdatedAt),
	}
}
