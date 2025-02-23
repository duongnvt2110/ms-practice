package auth

import (
	"auth-service/pkg/handler/http/auth/dto"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *authHandler) Register(c echo.Context) error {
	regBody := dto.RegisterRequestForm{}
	err := c.Bind(&regBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate
	err = c.Validate(regBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// Handing Bussines Login

	// Create oauthState cookie
	return c.JSON(http.StatusOK, nil)
}

func (h *authHandler) Login(c echo.Context) error {
	loginBody := dto.LoginRequestForm{}
	err := c.Bind(&loginBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// Validate
	err = c.Validate(loginBody)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	// Handing Bussines Login

	// Create oauthState cookie
	return c.JSON(http.StatusOK, nil)
}

func (h *authHandler) Logout(c echo.Context) error {
	// Handing Bussines Login

	// Create oauthState cookie
	return c.JSON(http.StatusOK, nil)
}
