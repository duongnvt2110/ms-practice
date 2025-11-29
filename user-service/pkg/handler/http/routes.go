package http_handler

import (
	"encoding/json"
	"log"
	"ms-practice/user-service/pkg/container"
	"ms-practice/user-service/pkg/handler/http/user"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func SetRoutes(r *mux.Router, c *container.Container) {
	userHandler := user.NewUserHandler(c.Cfg, *c.Usecase)
	userRoutes := r.PathPrefix("/v1/users").Subrouter()
	userRoutes.HandleFunc("/me", userHandler.GetMe).Methods("GET")
	userRoutes.HandleFunc("/{id:[0-9]+}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(10 * time.Second)
		log.Println("testGracefulShutdown job completed")
		result, _ := json.Marshal(map[string]interface{}{"status": "completed"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(result)
	}).Methods("GET")
}
