package clients

import (
	"context"
	"log"
	"ms-practice/auth-service/pkg/pb"

	"google.golang.org/grpc"
)

type UserClient struct {
	client pb.UserServiceClient
}

func NewUserClient(grpcAddr string) *UserClient {
	conn, err := grpc.Dial(grpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to user-service: %v", err)
	}

	return &UserClient{
		client: pb.NewUserServiceClient(conn),
	}
}

func (uc *UserClient) GetUser(ctx context.Context, id int32) (*pb.User, error) {
	resp, err := uc.client.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}
