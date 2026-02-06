package echo

import (
	"errors"
	"fmt"
	"ms-practice/pkg/errorsx"
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Send sends a JSON response.
func ResponseWithSuccess(c echo.Context, data interface{}) error {
	response := APIResponse{
		Message: "OK",
		Code:    fmt.Sprint(http.StatusOK),
		Status:  http.StatusOK,
		Data:    data,
	}

	return c.JSON(http.StatusOK, response)
}

func ResponseWithError(c echo.Context, errApp error) error {
	var err errorsx.AppError
	if !errors.As(errApp, &err) {
		err = errorsx.ErrInternalServerError.Wrap(errApp)
	}

	response := APIResponse{
		Message: err.PublicMessage(),
		Code:    err.GetErrCode(),
		Status:  err.GetHttpCode(),
	}

	return c.JSON(err.GetHttpCode(), response)
}
