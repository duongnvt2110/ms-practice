package user

import (
	"context"
	"log"
	"ms-practice/proto/gen"
	"ms-practice/user-service/pkg/models"
	"ms-practice/user-service/pkg/usecases"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserHandler struct {
	gen.UnimplementedUserServiceServer
	usecase usecases.UserUsecase
}

func NewUserHandler(uc usecases.UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) GetUser(ctx context.Context, req *gen.GetUserRequest) (*gen.GetUserResponse, error) {
	user, err := h.usecase.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &gen.GetUserResponse{
		User: &gen.User{
			Id:          int32(user.Id),
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Birthday:    user.Birthday,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}

func (h *UserHandler) CreateUser(ctx context.Context, req *gen.CreateUserRequest) (*gen.CreateUserResponse, error) {
	user := &models.User{
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Birthday:    req.Birthday,
		PhoneNumber: req.PhoneNumber,
	}

	id, err := h.usecase.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &gen.CreateUserResponse{Id: id}, nil
}

func (h *UserHandler) TestGracefulShutdown(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	time.Sleep(10 * time.Second)
	log.Println("testGracefulShutdown job completed")
	return &emptypb.Empty{}, nil
}
