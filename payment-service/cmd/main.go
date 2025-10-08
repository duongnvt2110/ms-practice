package main

import (
	"context"
	"os"
	"os/signal"
	"sync"

	"ms-practice/payment-service/pkg/container"
)

func main() {
	// test()
	c := container.InitializeContainer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	var wg sync.WaitGroup
	// Run HTTP Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		http_handler.StartHTTPServer(c, ctx)
	}()

	// Run GRPC Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		grpc_handler.StartGRPCUserServiceServer(c, ctx)
	}()
	<-ctx.Done()
	wg.Wait()
}
