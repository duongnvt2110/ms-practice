package main

import (
	"ms-practice/booking-service/pkg/container"
	http_handler "ms-practice/booking-service/pkg/handler/http"
)

func main() {
	// test()
	c := container.InitializeContainer()
	// Run HTTP Server
	http_handler.StartHTTPServer(c)
	select {}
}
