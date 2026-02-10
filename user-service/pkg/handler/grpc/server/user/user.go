package user

import (
	"context"
	"log"
	"ms-practice/proto/gen"
	"ms-practice/user-service/pkg/model"
	"ms-practice/user-service/pkg/usecase"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	gen.UnimplementedUserServiceServer
	usecase usecase.UserUC
}

func NewUserHandler(uc usecase.UserUC) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	user, err := h.usecase.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &gen.GetUserResponse{
		User: &gen.User{
			Id:           int32(user.Id),
			Email:        user.Email,
			Username:     user.Username,
			Avatar:       user.Avatar,
			Birthday:     user.Birthday,
			MobileNumber: user.MobileNumber,
		},
	}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	user := &model.User{
		Email:        req.Email,
		Username:     req.Username,
		Avatar:       req.Avatar,
		Birthday:     req.Birthday,
		MobileNumber: req.MobileNumber,
	}

	id, err := h.usecase.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &gen.CreateUserResponse{Id: id}, nil
}

func (h *UserHandler) DeleteUser(ctx context.Context, req *gen.DeleteUserRequest) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, h.usecase.DeleteUser(ctx, req.Id)
}

func (h *UserHandler) TestGracefulShutdown(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	time.Sleep(10 * time.Second)
	log.Println("testGracefulShutdown job completed")
	return &emptypb.Empty{}, nil
}
