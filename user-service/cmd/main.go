package main

import (
	"ms-practice/user-service/pkg/container"
	grpc_handler "ms-practice/user-service/pkg/handler/grpc/server"
	http_handler "ms-practice/user-service/pkg/handler/http"
)

func main() {
	// test()
	c := container.InitializeContainer()
	// Run HTTP Server
	http_handler.StartHTTPServer(c)
	// Run GRPC Server
	grpc_handler.StartGRPCUserServiceServer(c)

	select {}
}
