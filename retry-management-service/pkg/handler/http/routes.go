package http_handler

import (
	"net/http"

	"ms-practice/retry-management-service/pkg/container"
	admin_handler "ms-practice/retry-management-service/pkg/handler/http/admin"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, c *container.Container) {
	adminHandler, err := admin_handler.NewAdminHandler(c.DLQUC)
	if err != nil {
		panic(err)
	}

	r.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	}).Methods("GET")

	adminRoutes := r.PathPrefix("/admin").Subrouter()
	adminRoutes.HandleFunc("/dlq", adminHandler.List).Methods("GET")
	adminRoutes.HandleFunc("/dlq/{id:[0-9]+}", adminHandler.Detail).Methods("GET")
}
