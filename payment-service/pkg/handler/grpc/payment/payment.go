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
			Id:               int32(payment.Id),
			IdempotencyKey:   payment.IdempotencyKey,
			PaymentCode:      payment.PaymentCode,
			UserId:           int32(payment.UserId),
			BookingId:        int32(payment.BookingID),
			TransactionId:    payment.TransactionID,
			Amount:           int32(payment.Amount),
			Status:           payment.Status,
			Provider:         payment.Provider,
			PaidAt:           toProtoTimestamp(payment.PaidAt),
			CreatedAt:        timestamppb.New(payment.CreatedAt),
			UpdatedAt:        timestamppb.New(payment.UpdatedAt),
			PaymentHistories: toProtoPaymentHistories(payment.PaymentHistories),
		},
	}, nil
}

func toProtoPaymentHistories(histories []model.PaymentHistory) []*gen.PaymentHistory {
	if len(histories) == 0 {
		return nil
	}
	result := make([]*gen.PaymentHistory, 0, len(histories))
	for _, history := range histories {
		result = append(result, &gen.PaymentHistory{
			Id:        int32(history.Id),
			PaymentId: int32(history.PaymentID),
			Status:    history.Status,
			Logs:      history.Logs,
			PaidAt:    formatTime(history.PaidAt),
			CreatedAt: timestamppb.New(history.CreatedAt),
			UpdatedAt: timestamppb.New(history.UpdatedAt),
		})
	}
	return result
}

func toProtoTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}
	return timestamppb.New(*t)
}

func formatTime(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format(time.RFC3339)
}
