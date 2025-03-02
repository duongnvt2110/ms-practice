package auth

import (
	"net/http"

	"ms-practice/auth-service/pkg/handler/http/auth/dto"
	"ms-practice/auth-service/pkg/models"

	"github.com/labstack/echo/v4"
)

// Register endpoint
func (h *AuthHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(dto.RegisterRequestForm)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate
	if err := h.validate.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Call Usecase
	authProfileInfo := &models.AuthProfile{
		Password: req.Password,
		Email:    req.Email,
	}

	userInfo := &models.User{
		Birthday:    req.BirthDay,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
	}

	err := h.authProfileUC.Register(ctx, authProfileInfo, userInfo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

// Login endpoint
func (h *AuthHandler) Login(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(dto.LoginRequestForm)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Validate
	if err := h.validate.Struct(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	// Call Usecase
	token, err := h.authProfileUC.Login(ctx, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, token)
}

// Logout endpoint
func (h *AuthHandler) Logout(c echo.Context) error {
	ctx := c.Request().Context()
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusBadRequest, "token required")
	}

	// Call Usecase
	err := h.authProfileUC.Logout(ctx, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, "logout successful")
}
