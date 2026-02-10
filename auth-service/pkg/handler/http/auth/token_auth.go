package auth

import (
	"ms-practice/auth-service/pkg/handler/http/auth/dto"
	"ms-practice/auth-service/pkg/models"
	"ms-practice/auth-service/pkg/utils/apperr"
	resp "ms-practice/pkg/http/echo"

	"github.com/labstack/echo/v4"
)

// Register endpoint
func (h *AuthHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := &dto.RegisterRequestForm{}
	if err := c.Bind(req); err != nil {
		return resp.ResponseWithError(c, err)
	}

	// Validate
	if err := h.validate.Struct(req); err != nil {
		return resp.ResponseWithError(c, err)
	}

	// Call Usecase
	authProfileInfo := &models.AuthProfile{
		Password: req.Password,
		Email:    req.Email,
		Username: req.Username,
	}

	userInfo := &models.User{
		Email:        req.Email,
		Birthday:     req.Birthday,
		Username:     req.Username,
		Avatar:       req.Avatar,
		MobileNumber: req.MobileNumber,
	}

	err := h.authProfileUC.Register(ctx, authProfileInfo, userInfo)
	if err != nil {
		return resp.ResponseWithError(c, err)
	}

	return resp.ResponseWithSuccess(c, nil)
}

// Login endpoint
func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := &dto.LoginRequestForm{}
	if err := c.Bind(req); err != nil {
		return resp.ResponseWithError(c, err)
	}

	// Validate
	if err := h.validate.Struct(req); err != nil {
		return resp.ResponseWithError(c, err)
	}

	// Call Usecase
	token, err := h.authProfileUC.Login(ctx, req.Email, req.Password)
	if err != nil {
		return resp.ResponseWithError(c, err)
	}

	return resp.ResponseWithSuccess(c, token)
}

// RefreshToken endpoint
func (h *AuthHandler) RefreshToken(c echo.Context) error {
	ctx := c.Request().Context()
	req := &dto.RefreshTokenRequest{}
	if err := c.Bind(req); err != nil {
		return resp.ResponseWithError(c, err)
	}

	if err := h.validate.Struct(req); err != nil {
		return resp.ResponseWithError(c, err)
	}

	tokenPair, err := h.authProfileUC.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return resp.ResponseWithError(c, err)
	}

	return resp.ResponseWithSuccess(c, tokenPair)
}

// Logout endpoint
func (h *AuthHandler) Logout(c echo.Context) error {
	ctx := c.Request().Context()
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return resp.ResponseWithError(c, apperr.ErrTokenRequired)
	}
	authProfileID := c.Get("auth_profile_id").(int)

	// Call Usecase
	err := h.authProfileUC.Logout(ctx, authProfileID, token)
	if err != nil {
		return resp.ResponseWithError(c, err)
	}

	return resp.ResponseWithSuccess(c, "logout successful")
}
