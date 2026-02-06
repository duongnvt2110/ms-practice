package grpc

import (
	"context"
	"fmt"
	"log"
	"ms-practice/proto/gen"
	"ms-practice/user-service/pkg/container"
	"ms-practice/user-service/pkg/handler/grpc/server/user"
	"net"
	"time"

	"google.golang.org/grpc"
)

func StartGRPCUserServiceServer(ctx context.Context, c *container.Container) {
	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Cfg.GrpcUserSvc.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	userHandler := user.NewUserHandler(c.Usecase.UserUC)

	grpcServer := grpc.NewServer()
	gen.RegisterUserServiceServer(grpcServer, userHandler)

	errCh := make(chan error)

	// Run gRPC in a separate goroutine
	go func() {
		defer close(errCh)
		log.Printf("UserService gRPC Server is running on port %s...", c.Cfg.GrpcUserSvc.Port)
		if err := grpcServer.Serve(lis); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		// Gracefull shutdown
		gracefullShutdown(grpcServer)
		return
	case err := <-errCh:
		log.Printf("Server crash by %s:", err)
		return
	}
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
