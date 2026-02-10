package main

import (
	"context"
	"ms-practice/event-service/pkg/container"
	http_handler "ms-practice/event-service/pkg/handler/http"
	"os"
	"os/signal"
)

func main() {
	c := container.InitializeContainer()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		http_handler.StartHTTPServer(c, ctx)
	}()

	<-ctx.Done()
}
