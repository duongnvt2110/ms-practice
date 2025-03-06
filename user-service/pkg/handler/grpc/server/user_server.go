package grpc

import (
	"fmt"
	"log"
	"ms-practice/proto/gen"
	"ms-practice/user-service/pkg/container"
	"ms-practice/user-service/pkg/handler/grpc/server/user"
	"net"
	"os"

	"google.golang.org/grpc"
)

func StartGRPCUserServiceServer(c *container.Container) {
	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	userHandler := user.NewUserHandler(c.UserUc)

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
	}()
	// gracefullShutdown(srv)
	// log.Fatal("exit", <-errs)
}
