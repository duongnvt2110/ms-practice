package main

import (
	"booking-service/pkg/container"
	http_handler "booking-service/pkg/handler/http"
)

func main() {
	// test()
	c := container.InitializeContainer()
	// Run HTTP Server
	http_handler.StartHTTPServer(c)
	select {}
}
