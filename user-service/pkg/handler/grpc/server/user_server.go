package grpc

import (
	"context"
	"fmt"
	"log"
	"ms-practice/proto/gen"
	"ms-practice/user-service/pkg/container"
	"ms-practice/user-service/pkg/handler/grpc/server/user"
	"net"
	"os"
	"time"

	"google.golang.org/grpc"
)

func StartGRPCUserServiceServer(c *container.Container, ctx context.Context) {
	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	userHandler := user.NewUserHandler(c.Usecase.UserUC)

	grpcServer := grpc.NewServer()
	gen.RegisterUserServiceServer(grpcServer, userHandler)

	logChannel := make(chan string)

	// Run gRPC in a separate goroutine
	go func() {
		logChannel <- fmt.Sprintf("UserService gRPC Server is running on port %s...", c.Cfg.GRPC.Port)
		if err := grpcServer.Serve(lis); err != nil {
			logChannel <- fmt.Sprintf("Failed to serve: %v", err)
			os.Exit(1)
		}
	}()

	go func() {
		for msg := range logChannel {
			fmt.Println(msg)
		}
		close(logChannel)
	}()
	<-ctx.Done()
	gracefullShutdown(grpcServer)
}

func gracefullShutdown(srv *grpc.Server) {
	log.Println("Shutting down Grpc server...")
	// Implement graceful shutdown.
	timer := time.AfterFunc(10*time.Second, func() {
		log.Println("Server couldn't stop gracefully in time. Doing force stop.")
		srv.Stop()
	})
	defer timer.Stop()
	srv.GracefulStop()
	log.Println("Grpc server exited gracefully")
}
