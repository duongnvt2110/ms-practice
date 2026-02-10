package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	paymentusecase "ms-practice/payment-service/pkg/usecase/payment"
	kafka_booking "ms-practice/payment-service/pkg/util/kafka"
	"ms-practice/pkg/event"

	"github.com/segmentio/kafka-go"
)

type PaymentConsumer struct {
	message   *kafka_booking.BookingMessaging
	paymentUC paymentusecase.PaymentUsecaseInterface
}

func NewPaymentConsumer(message *kafka_booking.BookingMessaging, paymentUC paymentusecase.PaymentUsecaseInterface) *PaymentConsumer {
	return &PaymentConsumer{message: message, paymentUC: paymentUC}
}

func (p *PaymentConsumer) Start(ctx context.Context) error {
	consumer, ok := p.message.Consumers[event.BookingTopicName]
	if !ok || consumer == nil {
		return nil
	}
	return consumer.Consume(ctx, p.handle)
}

func (p *PaymentConsumer) handle(k kafka.Message) error {
	var evt event.BookingPayload
	if err := json.Unmarshal(k.Value, &evt); err != nil {
		log.Printf("failed to unmarshal booking event: %v", err)
		return err
	}
	return errors.ErrUnsupported

	ctx := context.Background()
	payment, err := p.paymentUC.ProcessPayment(ctx, evt)
	payload := event.PaymentPayload{
		OrderID: evt.OrderID,
		UserID:  evt.UserID,
		Amount:  evt.Amount,
	}
	if err != nil {
		payload.EventType = event.PaymentFailed
		payload.Message = err.Error()
	} else {
		payload.EventType = event.PaymentSucceeded
		payload.PaymentID = payment.Id
	}
	b, marshalErr := json.Marshal(payload)
	if marshalErr != nil {
		log.Printf("failed to marshal payment payload: %v", marshalErr)
		return marshalErr
	}
	producer, ok := p.message.Producers[event.PaymentTopicName]
	if !ok || producer == nil {
		return errors.New("Kafka Unavaiable")
	}
	if publishErr := producer.Publish(ctx, nil, b); publishErr != nil {
		log.Printf("failed to publish payment event: %v", publishErr)
		return publishErr
	}
	return err
}
