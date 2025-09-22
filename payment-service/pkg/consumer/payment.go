package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"

	"payment-service/pkg/events"
	kafka_client "payment-service/pkg/utils/kafka"
)

type PaymentConsumer struct {
	kafka kafka_client.KafkaClient
}

func NewPaymentConsumer(k kafka_client.KafkaClient) *PaymentConsumer {
	return &PaymentConsumer{kafka: k}
}

func (p *PaymentConsumer) Start(ctx context.Context) error {
	return p.kafka.SetReaderTopic(events.TopicOrderEvents, "payment-service").Consume(ctx, p.handle)
}

func (p *PaymentConsumer) handle(m kafka.Message) {
	var evt events.OrderPlacedEvent
	if err := json.Unmarshal(m.Value, &evt); err != nil {
		log.Printf("failed to unmarshal event: %v", err)
		return
	}
	processed := events.PaymentProcessedEvent{
		OrderID:   evt.OrderID,
		PaymentID: uuid.NewString(),
		Success:   true,
	}
	b, err := json.Marshal(processed)
	if err != nil {
		log.Printf("failed to marshal payment event: %v", err)
		return
	}
	if err := p.kafka.SetWriterTopic(events.TopicPaymentEvents).Publish(context.Background(), nil, b); err != nil {
		log.Printf("failed to publish payment event: %v", err)
	}
}
