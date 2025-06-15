package main

import (
	"context"
	"log"

	"payment-service/pkg/consumer"
	"payment-service/pkg/container"
)

func main() {
	c := container.InitializeContainer()
	ctx := context.Background()
	if err := consumer.NewPaymentConsumer(c.Kafka).Start(ctx); err != nil && err != context.Canceled {
		log.Fatal(err)
	}
}
