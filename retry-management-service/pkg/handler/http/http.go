package http_handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"ms-practice/retry-management-service/pkg/container"

	"github.com/gorilla/mux"
)

func StartHTTPServer(ctx context.Context, c *container.Container) {
	r := mux.NewRouter()

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", c.Cfg.App.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      r,
	}

	RegisterRoutes(r, c)

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		log.Printf("Server is running on http://%s:%s", c.Cfg.App.Host, c.Cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		gracefulShutdown(srv)
		return
	case err := <-errCh:
		log.Printf("Server crash by %v", err)
		return
	}
}

func gracefulShutdown(srv *http.Server) {
	log.Println("Shutting down HTTP server...")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed %v", err)
	}
	log.Println("Server exited gracefully")
}
