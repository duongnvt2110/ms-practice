package http_middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetMiddleware(r *mux.Router) {
	r.Use(loggingMiddleware)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		w.Header().Set("Content-Type", "application/json")
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
