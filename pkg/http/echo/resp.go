package echo

import (
	"fmt"
	"net/http"

	apperror "ms-practice/auth-service/pkg/utils/app_error"

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
	err, ok := errApp.(apperror.AppError)
	if !ok {
		err = apperror.ErrInternalServer.Wrap(err)
	}

	response := APIResponse{
		Message: err.Error(),
		Code:    err.GetErrCode(),
		Status:  err.GetHttpCode(),
	}

	return c.JSON(err.GetHttpCode(), response)
}
