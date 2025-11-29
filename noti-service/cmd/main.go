package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"

	"ms-practice/noti-service/pkg/consumer"
	"ms-practice/noti-service/pkg/container"
	httpHandler "ms-practice/noti-service/pkg/handler/http"
)

func main() {
	c := container.InitializeContainer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup

	// HTTP server
	wg.Add(1)
	go func() {
		defer wg.Done()
		httpHandler.StartHTTPServer(ctx, c)
	}()

	eventConsumer := &consumer.EventConsumer{Messaging: c.Messaging, Usecase: c.Usecase}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := eventConsumer.StartPaymentConsumer(ctx); err != nil {
			log.Printf("payment consumer stopped: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := eventConsumer.StartTicketConsumer(ctx); err != nil {
			log.Printf("ticket consumer stopped: %v", err)
		}
	}()

	for _, w := range c.Workers {
		wg.Add(1)
		worker := w
		go func() {
			defer wg.Done()
			worker.Start(ctx)
		}()
	}

	<-ctx.Done()
	wg.Wait()
}
