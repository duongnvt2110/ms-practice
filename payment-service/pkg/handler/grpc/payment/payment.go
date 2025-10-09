package payment

import (
	"context"
	"ms-practice/payment-service/pkg/usecase/payment"
	"ms-practice/proto/gen"
)

type PaymentHandler struct {
	gen.UnsafePaymentServiceServer
	paymentUC payment.PaymentUsecaseInterface
}

func NewPaymentGrpcHandler(paymentUC payment.PaymentUsecaseInterface) *PaymentHandler {
	return &PaymentHandler{paymentUC: paymentUC}
}

func (h *PaymentHandler) GetPayment(ctx context.Context, req *gen.GetPaymentRequest) (*gen.GetPaymentResponse, error) {
	payment, err := h.paymentUC.GetPayment(ctx, req.Id, req.PaymentId)
	if err != nil {
		return nil, err
	}

	return &gen.GetPaymentResponse{
		Payment: &gen.Payment{
			Id:        int32(payment.Id),
			UserId:    int32(payment.UserId),
			BookingId: int32(payment.BookingID),
			Quantity: int32(payment.Quantity),
			Prices: float32(payment.Prices),
			CreatedAt: payment.CreatedAt,
			UpdatedAt: payment.UpdatedAt,
			PaymentHistory: &gen.PaymentHistory{},
		},
	}, nil
}

    int32 id = 1;
    int32 user_id = 2;
    int32 booking_id = 3;
    int32 quantity = 4;
    float prices = 5;
    string method = 6;
    google.protobuf.Timestamp created_at = 7;
    google.protobuf.Timestamp updated_at = 8;
    PaymentHistory payment_history = 9;