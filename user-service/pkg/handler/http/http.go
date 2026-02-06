package http_handler

import (
	"context"
	"fmt"
	"log"
	"ms-practice/user-service/pkg/container"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartHTTPServer(ctx context.Context, c *container.Container) {
	h := mux.NewRouter()
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", c.Cfg.App.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h,
	}

	RegisterRoutes(h, c)
	RegisterMiddleware(h, c.Usecase.AuthUC)

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
		// Gracefull shutdown
		gracefullShutdown(srv)
		return
	case err := <-errCh:
		log.Printf("Server crash by %s:", err)
		return
	}
}

func gracefullShutdown(srv *http.Server) {
	log.Println("Shutting down HTTP server...")
	// Implement graceful shutdown.
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("Server shutdown failed %s:", err)
	}
	log.Println("Server exited gracefully")
}
