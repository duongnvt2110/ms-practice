package consumer

import (
	"context"
	"encoding/json"
	"log"

	"ms-practice/booking-service/pkg/model"
	"ms-practice/booking-service/pkg/usecase/booking"
	bookingkafka "ms-practice/booking-service/pkg/util/kafka"
	"ms-practice/pkg/event"

	kafkago "github.com/segmentio/kafka-go"
)

type PaymentConsumer struct {
	messaging *bookingkafka.BookingMessaging
	bookingUC booking.BookingUsecase
}

func NewPaymentConsumer(bookingMessaging *bookingkafka.BookingMessaging, bookingUC booking.BookingUsecase) *PaymentConsumer {
	return &PaymentConsumer{
		messaging: bookingMessaging,
		bookingUC: bookingUC,
	}
}

func (p *PaymentConsumer) Start(ctx context.Context) error {
	consumer, ok := p.messaging.Consumers[event.PaymentTopicName]
	if !ok || consumer == nil {
		return nil
	}
	return consumer.Consume(ctx, func(msg kafkago.Message) {
		p.handle(ctx, msg)
	})
}

func (p *PaymentConsumer) handle(ctx context.Context, k kafkago.Message) {
	var payload event.PaymentPayload
	if err := json.Unmarshal(k.Value, &payload); err != nil {
		log.Printf("failed to unmarshal payment payload: %v", err)
		return
	}

	var status string
	switch payload.EventType {
	case event.PaymentSucceeded:
		status = model.BookingStatusConfirmed
	case event.PaymentFailed:
		status = model.BookingStatusFailed
	default:
		log.Printf("ignored payment event with unknown type %q for order %d", payload.EventType, payload.OrderID)
		return
	}

	updateCtx := ctx
	if updateCtx == nil {
		updateCtx = context.Background()
	}
	if err := p.bookingUC.UpdateBookingStatus(updateCtx, payload.OrderID, status); err != nil {
		log.Printf("failed to update booking %d status: %v", payload.OrderID, err)
	}
}
