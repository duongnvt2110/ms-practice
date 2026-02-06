package grpc

import (
	"context"
	"fmt"
	"log"
	"ms-practice/auth-service/pkg/container"
	"ms-practice/auth-service/pkg/handler/grpc/auth"
	"ms-practice/proto/gen"
	"net"
	"time"

	"google.golang.org/grpc"
)

func StartGrpcServer(ctx context.Context, c *container.Container) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", c.Cfg.GrpcAuthSvc.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authHandler := auth.NewAuthHandler(c.Usecase.AuthProfileUC)
	grpcServer := grpc.NewServer()
	gen.RegisterAuthServiceServer(grpcServer, authHandler)

	errCh := make(chan error)

	go func() {
		defer close(errCh)
		log.Printf("AuthService gRPC Server is running on port %v", c.Cfg.GrpcAuthSvc.Port)
		if err := grpcServer.Serve(lis); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		gracefullShutdown(grpcServer)
		return
	case err := <-errCh:
		log.Printf("server crashed by %v", err)
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
