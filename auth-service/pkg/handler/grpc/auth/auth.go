package auth

import (
	"context"
	"ms-practice/auth-service/pkg/usecase"
	"ms-practice/pkg/errorsx"
	"ms-practice/proto/gen"
)

type AuthHandler struct {
	gen.UnimplementedAuthServiceServer
	authProfileUC usecase.AuthProfileUC
}

func NewAuthHandler(authProfileUC usecase.AuthProfileUC) *AuthHandler {
	return &AuthHandler{
		authProfileUC: authProfileUC,
	}
}

func (h *AuthHandler) ValidateToken(ctx context.Context, authRequest *gen.ValAuthReq) (*gen.ValAuthResp, error) {
	authClaims, err := h.authProfileUC.ValidateToken(authRequest.AuthToken)
	if err != nil {
		return nil, errorsx.ToStatus(err)
	}
	authResp := &gen.ValAuthResp{
		AuthProfileId: int32(authClaims.AuthProfileId),
	}
	return authResp, nil
}
