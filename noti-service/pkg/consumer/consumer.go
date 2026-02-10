package consumer

import (
	"context"
	"encoding/json"
	"log"

	"ms-practice/noti-service/pkg/usecase"
	nkafka "ms-practice/noti-service/pkg/util/kafka"
	"ms-practice/pkg/event"

	kafkago "github.com/segmentio/kafka-go"
)

type EventConsumer struct {
	Messaging *nkafka.NotificationMessaging
	Usecase   usecase.NotificationUsecase
}

func (c *EventConsumer) StartPaymentConsumer(ctx context.Context) error {
	return c.start(ctx, event.PaymentTopicName, c.handlePayment)
}

func (c *EventConsumer) StartTicketConsumer(ctx context.Context) error {
	return c.start(ctx, event.TicketTopicName, c.handleTicket)
}

func (c *EventConsumer) start(ctx context.Context, topic string, handler func(context.Context, kafkago.Message) error) error {
	if c.Messaging == nil {
		return nil
	}
	client, ok := c.Messaging.Consumers[topic]
	if !ok || client == nil {
		return nil
	}
	return client.Consume(ctx, func(msg kafkago.Message) error {
		if err := handler(ctx, msg); err != nil {
			return err
		}
		return nil
	})
}

func (c *EventConsumer) handlePayment(ctx context.Context, msg kafkago.Message) error {
	var payload event.PaymentPayload
	if err := json.Unmarshal(msg.Value, &payload); err != nil {
		log.Printf("noti payment consumer failed to unmarshal: %v", err)
		return err
	}
	if err := c.Usecase.HandlePaymentEvent(ctx, payload); err != nil {
		log.Printf("noti payment consumer failed to handle event: %v", err)
		return err
	}
	return nil
}

func (c *EventConsumer) handleTicket(ctx context.Context, msg kafkago.Message) error {
	var payload event.TicketPayload
	if err := json.Unmarshal(msg.Value, &payload); err != nil {
		log.Printf("noti ticket consumer failed to unmarshal: %v", err)
		return err
	}
	if err := c.Usecase.HandleTicketEvent(ctx, payload); err != nil {
		log.Printf("noti ticket consumer failed to handle event: %v", err)
		return err
	}
	return nil
}
