package main

import (
	"context"
	"log"

	"ms-practice/payment-service/pkg/consumer"
	"ms-practice/payment-service/pkg/container"
	"ms-practice/payment-service/pkg/util/kafka"
	sharedKafka "ms-practice/pkg/kafka"
)

func main() {
	c := container.InitializeContainer()
	ctx := context.Background()
	kafkaClient := sharedKafka.NewKafkaClient(c.Cfg.App.Kafka)
	bookingMessaging := kafka.NewBookingKafkaClient(kafkaClient)
	log.Println("Booking Consumer Service is running")
	if err := consumer.NewPaymentConsumer(bookingMessaging).Start(ctx); err != nil && err != context.Canceled {
		log.Fatal(err)
	}
}
