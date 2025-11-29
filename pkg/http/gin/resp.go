package gin

import (
	"fmt"
	apperror "ms-practice/booking-service/pkg/util/app_error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Send sends a JSON response.
func ResponseWithSuccess(c *gin.Context, data interface{}) {
	response := APIResponse{
		Message: "OK",
		Code:    fmt.Sprint(http.StatusOK),
		Status:  http.StatusOK,
		Data:    data,
	}

	c.JSON(http.StatusOK, response)
}

func ResponseWithError(c *gin.Context, errApp error) {
	err, ok := errApp.(apperror.AppError)
	if !ok {
		err = apperror.ErrInternalServer.Wrap(err)
	}

	response := APIResponse{
		Message: err.Error(),
		Code:    err.GetErrCode(),
		Status:  err.GetHttpCode(),
	}

	c.JSON(err.GetHttpCode(), response)
}
