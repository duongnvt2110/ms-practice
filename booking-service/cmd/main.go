package main

import (
	"context"
	"log"
	"ms-practice/booking-service/pkg/consumer"
	"ms-practice/booking-service/pkg/container"
	http_handler "ms-practice/booking-service/pkg/handler/http"
)

func main() {
	c := container.InitializeContainer()

	ctx := context.Background()
	go func() {
		err := consumer.NewPaymentConsumer(c.BookingMessaging, c.Usecases.BookingUC).Start(ctx)
		if err != nil && err != context.Canceled {
			log.Fatalf("payment consumer stopped: %v", err)
		}
	}()

	http_handler.StartHTTPServer(c)
	select {}
}
