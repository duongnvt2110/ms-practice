package main

import (
	"context"
	"log"
	"ms-practice/booking-service/pkg/consumer"
	"ms-practice/booking-service/pkg/container"
	http_handler "ms-practice/booking-service/pkg/handler/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	c := container.InitializeContainer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		err := consumer.NewPaymentConsumer(c.BookingMessaging, c.Usecases.BookingUC).Start(gCtx)
		if err != nil && err != context.Canceled {
			return err
		}
		return nil
	})

	g.Go(func() error {
		http_handler.StartHTTPServer(gCtx, c)
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("service stopped with error: %v", err)
		return
	}

	log.Printf("service stopped gracefully")
}
