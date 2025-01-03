package main

import (
	"user-service/pkg/container"
	http_handler "user-service/pkg/handler/http"
)

func main() {
	// test()
	c := container.InitializeContainer()
	// Run HTTP Server
	http_handler.StartHTTPServer(c)
	select {}
}
