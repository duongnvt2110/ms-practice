package main

import (
	"context"
	"ms-practice/event-service/pkg/container"
	http_handler "ms-practice/event-service/pkg/handler/http"
	"os"
	"os/signal"
	"sync"
)

func main() {
	c := container.InitializeContainer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		http_handler.StartHTTPServer(c, ctx)
	}()

	<-ctx.Done()
	wg.Wait()
}
