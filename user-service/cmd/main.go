package main

import (
	"context"
	"ms-practice/user-service/pkg/container"
	grpc_handler "ms-practice/user-service/pkg/handler/grpc/server"
	http_handler "ms-practice/user-service/pkg/handler/http"
	"os"
	"os/signal"
	"sync"
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
