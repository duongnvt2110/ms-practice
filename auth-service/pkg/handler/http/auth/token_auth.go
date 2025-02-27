package auth

import (
	"auth-service/pkg/handler/http/auth/dto"
	"auth-service/pkg/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Register endpoint
func (h *AuthHandler) Register(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(dto.RegisterRequestForm)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// Call Usecase
	authProfileInfo := &models.AuthProfile{
		Password: req.Password,
		Email:    req.Email,
	}

	userInfo := &models.User{
		Birthday:  req.BirthDay,
		FirstName: req.FirstName,
		LastName:  req.LastName,
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
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate
	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
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
