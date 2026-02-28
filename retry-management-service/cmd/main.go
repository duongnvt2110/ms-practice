package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"ms-practice/retry-management-service/pkg/consumer"
	"ms-practice/retry-management-service/pkg/container"
	http_handler "ms-practice/retry-management-service/pkg/handler/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	c := container.InitializeContainer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		http_handler.StartHTTPServer(gCtx, c)
		return nil
	})

	g.Go(func() error {
		return consumer.StartDLQConsumer(gCtx, c.Cfg, c.DLQUC)
	})

	if err := g.Wait(); err != nil {
		log.Printf("service stopped with error: %v", err)
		return
	}

	log.Printf("service stopped gracefully")
}
