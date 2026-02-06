package main

import (
	"context"
	"log"
	"ms-practice/user-service/pkg/container"
	grpc_handler "ms-practice/user-service/pkg/handler/grpc/server"
	http_handler "ms-practice/user-service/pkg/handler/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func main() {
	// test()
	c := container.InitializeContainer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		http_handler.StartHTTPServer(gCtx, c)
		return nil
	})
	// Run GRPC Server
	g.Go(func() error {
		grpc_handler.StartGRPCUserServiceServer(gCtx, c)
		return nil
	})

	if err := g.Wait(); err != nil {
		log.Printf("service stopped with error: %v", err)
		return
	}

	log.Printf("service stopped gracefully")
}
