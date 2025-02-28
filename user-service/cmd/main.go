package main

import (
	"fmt"
	"log"
	"net"
	"user-service/pkg/container"
	http_handler "user-service/pkg/handler/http"

	"google.golang.org/grpc"
)

func main() {
	// test()
	c := container.InitializeContainer()
	// Run HTTP Server
	http_handler.StartHTTPServer(c)
	
	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	fmt.Println("UserService gRPC Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	select {}
}
