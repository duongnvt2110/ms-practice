package usecases

import (
	"context"
	"fmt"
	"log"
	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/models"
	"ms-practice/proto/gen"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserGrpcClient interface {
	GetUser(ctx context.Context, id int32) (*gen.User, error)
	CreateUser(ctx context.Context, user *models.User) (*gen.CreateUserResponse, error)
	DeleteUser(ctx context.Context, userID int32) (*emptypb.Empty, error)
}

type userGrpcClient struct {
	client gen.UserServiceClient
}

var _ UserGrpcClient = (*userGrpcClient)(nil)

func NewUserGrpcClient(cfg *config.Config) UserGrpcClient {
	addr := fmt.Sprintf("%s:%s", cfg.GRPC.UserServiceHost, cfg.GRPC.UserServicePort)
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
		Email:        user.Email,
		Username:     user.Username,
		Avatar:       user.Avatar,
		Birthday:     user.Birthday,
		MobileNumber: user.MobileNumber,
	}
	resp, err := uc.client.CreateUser(ctx, userProto)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (uc *userGrpcClient) DeleteUser(ctx context.Context, userID int32) (*emptypb.Empty, error) {
	userProto := &gen.DeleteUserRequest{
		Id: userID,
	}
	return uc.client.DeleteUser(ctx, userProto)
}
