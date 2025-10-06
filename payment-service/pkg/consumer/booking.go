package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/segmentio/kafka-go"

	event "ms-practice/payment-service/pkg/event"
	kafka_client "ms-practice/payment-service/pkg/util/kafka"
)

type PaymentConsumer struct {
	kafka kafka_client.KafkaClient
}

func NewPaymentConsumer(k kafka_client.KafkaClient) *PaymentConsumer {
	return &PaymentConsumer{kafka: k}
}

func (p *PaymentConsumer) Start(ctx context.Context) error {
	return p.kafka.SetReaderTopic(event.BookingTopic, "booking.consumer").Consume(ctx, p.handle)
}

func (p *PaymentConsumer) handle(m kafka.Message) {
	var evt event.BookingOrdered
	if err := json.Unmarshal(m.Value, &evt); err != nil {
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
