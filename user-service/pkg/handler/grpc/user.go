package handler

import (
	"context"
	"ms-practice/user-service/pkg/pb"
	"ms-practice/user-service/pkg/usecase"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	usecase usecase.UserUsecase
}

func NewUserHandler(uc usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: uc}
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := h.usecase.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:          user.Id,
			Email:       user.Email,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Birthday:    user.Birthday,
			PhoneNumber: user.PhoneNumber,
		},
	}, nil
}
