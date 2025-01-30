package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	apperror "user-service/pkg/utils/app_error"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Send sends a JSON response.
func ResponseWithSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := APIResponse{
		Message: "OK",
		Code:    fmt.Sprint(http.StatusOK),
		Status:  http.StatusOK,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func ResponseWithError(w http.ResponseWriter, errApp error) {
	err, ok := errApp.(apperror.AppError)
	if !ok {
		err = apperror.ErrInternalServer.Wrap(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.GetHttpCode())

	response := APIResponse{
		Message: err.Error(),
		Code:    err.GetErrCode(),
		Status:  err.GetHttpCode(),
	}

	json.NewEncoder(w).Encode(response)
}
