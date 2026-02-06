package http_handler

import (
	"log"
	response "ms-practice/pkg/http/mux"
	"ms-practice/user-service/pkg/usecase"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func RegisterMiddleware(r *mux.Router, uc usecase.AuthUC) {
	r.Use(loggingMiddleware)
	r.Use(authMiddleware(uc))
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

func authMiddleware(authUC usecase.AuthUC) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Do stuff here
			auth := r.Header.Get("Authorization")
			token := strings.TrimPrefix(auth, "Bearer")
			token = strings.TrimSpace(token)

			authProfileId, err := authUC.ValidateToken(r.Context(), token)
			if err != nil {
				spew.Dump(err)
				response.ResponseWithError(w, err)
				return
			}
			context.Set(r, "auth_profile_id", authProfileId)
			next.ServeHTTP(w, r)
		})
	}
}
