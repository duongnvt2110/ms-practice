package grpc

import (
	"context"
	"fmt"
	"log"
	"ms-practice/payment-service/pkg/container"
	"ms-practice/payment-service/pkg/handler/grpc/payment"
	"ms-practice/proto/gen"
	"net"
	"time"

	"google.golang.org/grpc"
)

func StartGRPCUserServiceServer(c *container.Container, ctx context.Context) {
	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", c.Cfg.GrpcPaymentSvc.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	paymentGrpcHandler := payment.NewPaymentGrpcHandler(c.Usecases.PaymentUC)

	grpcServer := grpc.NewServer()
	gen.RegisterPaymentServiceServer(grpcServer, paymentGrpcHandler)

	errCh := make(chan error)

	// Run gRPC in a separate goroutine
	go func() {
		log.Printf("UserService gRPC Server is running on port %s...", c.Cfg.GrpcPaymentSvc.Port)
		defer close(errCh)
		if err := grpcServer.Serve(lis); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		gracefullShutdown(grpcServer)
		return
	case err, ok := <-errCh:
		if !ok {
			return
		}
		fmt.Printf("Failed to serve: %v", err)
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
