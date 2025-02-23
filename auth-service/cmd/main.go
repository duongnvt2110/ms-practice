package main

import (
	"auth-service/pkg/container"
	http_handler "auth-service/pkg/handler/http"
)

func main() {
	c := container.InitializeContainer()
	// // Run HTTP Server
	http_handler.StartHTTPServer(c)
	// select {}
}
