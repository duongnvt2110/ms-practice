package main

import (
	"context"
	"log"

	"ms-practice/payment-service/pkg/consumer"
	"ms-practice/payment-service/pkg/container"
	"ms-practice/payment-service/pkg/util/kafka"
)

func main() {
	c := container.InitializeContainer()
	ctx := context.Background()
	bookingMessaging := kafka.NewBookingKafkaClient(c.Cfg.App.Kafka)
	log.Println("Booking Consumer Service is running")
	if err := consumer.NewPaymentConsumer(bookingMessaging, c.Usecases.PaymentUC).Start(ctx); err != nil && err != context.Canceled {
		log.Fatal(err)
	}
}
