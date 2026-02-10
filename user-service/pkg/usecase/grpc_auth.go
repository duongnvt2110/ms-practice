package usecase

import (
	"context"
	"fmt"
	"log"
	"ms-practice/proto/gen"
	"ms-practice/user-service/pkg/config"

	"google.golang.org/grpc"
)

type AuthUC interface {
	ValidateToken(ctx context.Context, authToken string) (*int32, error)
}
type authUC struct {
	authGRPCServer gen.AuthServiceClient
}

var _ AuthUC = (*authUC)(nil)

func NewAuthUC(cfg *config.Config) AuthUC {
	addr := fmt.Sprintf("%s:%s", cfg.GrpcAuthSvc.Host, cfg.GrpcAuthSvc.Port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Auth Service: %v", err)
	}
	log.Printf("Connected to GRPC Auth Service on %v", addr)

	return &authUC{
		authGRPCServer: gen.NewAuthServiceClient(conn),
	}
}

func (uc *authUC) ValidateToken(ctx context.Context, authToken string) (*int32, error) {
	resp, err := uc.authGRPCServer.ValidateToken(ctx, &gen.ValAuthReq{
		AuthToken: authToken,
	})
	if err != nil {
		return nil, err
	}
	return &resp.AuthProfileId, err
}
