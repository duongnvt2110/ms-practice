package usecases

import (
	"context"
	"fmt"
	"log"
	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/models"
	"ms-practice/proto/gen"

	"google.golang.org/grpc"
)

type UserGrpcClient interface {
	GetUser(ctx context.Context, id int32) (*gen.User, error)
	CreateUser(ctx context.Context, user *models.User) (*gen.CreateUserResponse, error)
}

type userGrpcClient struct {
	client gen.UserServiceClient
}

var _ UserGrpcClient = (*userGrpcClient)(nil)

func NewUserGrpcClient(cfg *config.Config) UserGrpcClient {
	addr := fmt.Sprintf("%s:%s", cfg.GRPC.UserHost, cfg.GRPC.UserPort)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}
	log.Println("Connected to GRPC user-service")

	return &userGrpcClient{
		client: gen.NewUserServiceClient(conn),
	}
}

func (uc *userGrpcClient) GetUser(ctx context.Context, id int32) (*gen.User, error) {
	resp, err := uc.client.GetUser(ctx, &gen.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}

func (uc *userGrpcClient) CreateUser(ctx context.Context, user *models.User) (*gen.CreateUserResponse, error) {
	userProto := &gen.CreateUserRequest{
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Birthday:    user.Birthday,
		PhoneNumber: user.PhoneNumber,
	}
	resp, err := uc.client.CreateUser(ctx, userProto)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
