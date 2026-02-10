package http_handler

import (
	"context"
	"fmt"
	"log"
	"ms-practice/event-service/pkg/container"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer(c *container.Container, ctx context.Context) {
	engine := gin.Default()

	RegisterRoutes(engine, c)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", c.Cfg.App.Port),
		Handler:      engine,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	errCh := make(chan error)

	go func() {
		defer close(errCh)
		log.Printf("Server is running on http://%s:%s", c.Cfg.App.Host, c.Cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case <-ctx.Done():
		gracefullShutdown(srv)
		return
	case err := <-errCh:
		log.Printf("Server crash by %s:", err)
		return
	}
}

func gracefullShutdown(srv *http.Server) {
	shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := srv.Shutdown(shutdownContext)
	if err != nil {
		log.Printf("Server shutdown failed %s:", err)
	}
	log.Println("Server exited gracefully")
}
