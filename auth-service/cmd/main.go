package main

import (
	http_handler "ms-practice/auth-service/pkg/handler/http"

	"ms-practice/auth-service/pkg/container"
)

func main() {
	c := container.InitializeContainer()
	// // Run HTTP Server
	http_handler.StartHTTPServer(c)
}
