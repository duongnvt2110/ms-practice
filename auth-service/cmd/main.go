package main

import (
	"context"
	"log"
	"ms-practice/auth-service/pkg/handler/grpc"
	http_handler "ms-practice/auth-service/pkg/handler/http"
	"os"
	"os/signal"

	"ms-practice/auth-service/pkg/container"

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
		grpc.StartGrpcServer(gCtx, c)
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("service stopped with error: %v", err)
		return
	}

	log.Printf("service stopped gracefully")
}
