package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"` // Optional metadata (e.g., for pagination)
}

// Send sends a JSON response.
func ResponseWithSuccess(w http.ResponseWriter, message string, data interface{}, meta interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := APIResponse{
		Message: message,
		Data:    data,
		Meta:    meta,
	}

	json.NewEncoder(w).Encode(response)
}

func ResponseWithError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(err.Error())
}
