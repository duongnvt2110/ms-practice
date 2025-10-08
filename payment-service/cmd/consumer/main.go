package main

import (
	"context"
	"log"

	"ms-practice/payment-service/pkg/consumer"
	"ms-practice/payment-service/pkg/container"
)

func main() {
	c := container.InitializeContainer()
	ctx := context.Background()
	if err := consumer.NewPaymentConsumer(c.Kafka).Start(ctx); err != nil && err != context.Canceled {
		log.Fatal(err)
	}
}
