package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/segmentio/kafka-go"

	event "ms-practice/payment-service/pkg/event"
	kafka_booking "ms-practice/payment-service/pkg/util/kafka"
)

type PaymentConsumer struct {
	message *kafka_booking.BookingMessaging
}

func NewPaymentConsumer(message *kafka_booking.BookingMessaging) *PaymentConsumer {
	return &PaymentConsumer{message: message}
}

func (p *PaymentConsumer) Start(ctx context.Context) error {
	return p.message.Consumers[event.BoookingConsumerName].Consume(ctx, p.handle)
}

func (p *PaymentConsumer) handle(k kafka.Message) {
	var evt event.BookingPayload
	if err := json.Unmarshal(k.Value, &evt); err != nil {
		log.Printf("failed to unmarshal event: %v", err)
		return
	}
	spew.Dump(evt)
	// processed := events.PaymentProcessedEvent{
	// 	OrderID:   evt.OrderID,
	// 	PaymentID: uuid.NewString(),
	// 	Success:   true,
	// }
	// b, err := json.Marshal(processed)
	// if err != nil {
	// 	log.Printf("failed to marshal payment event: %v", err)
	// 	return
	// }
	// if err := p.kafka.SetWriterTopic(events.TopicPaymentEvents).Publish(context.Background(), nil, b); err != nil {
	// 	log.Printf("failed to publish payment event: %v", err)
	// }
}
