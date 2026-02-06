package mux

import (
	"encoding/json"
	"fmt"
	"ms-practice/pkg/errorsx"
	"net/http"
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

func ResponseWithError(w http.ResponseWriter, err error) {
	var e errorsx.AppError
	if apperr, ok := err.(errorsx.AppError); ok {
		e = apperr
	} else if gerr, ok := errorsx.FromStatus(err); ok {
		e = gerr.(errorsx.AppError)
	} else {
		e = errorsx.ErrInternalServerError.Wrap(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.GetHttpCode())

	response := APIResponse{
		Message: e.Error(),
		Code:    e.GetErrCode(),
		Status:  e.GetHttpCode(),
	}

	json.NewEncoder(w).Encode(response)
}
