package http_handler

import (
	"context"
	"fmt"
	"log"
	"ms-practice/booking-service/pkg/container"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer(ctx context.Context, c *container.Container) {
	h := gin.Default()
	addr := resolveAddr(c)
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h,
	}
	RegisterRoutes(h, c.Cfg, c.Usecases)
	// http_middleware.SetMiddleware(h)

	errCh := make(chan error)
	go func() {
		defer close(errCh)
		log.Printf("Server is running on http://%s", addr)
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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown failed %s:", err)
	}
	log.Println("Server exited gracefully")
}

func resolveAddr(c *container.Container) string {
	host := c.Cfg.App.Host
	port := c.Cfg.App.Port
	if port == "" {
		port = "3000"
	}
	if host == "" {
		return fmt.Sprintf(":%s", port)
	}
	return fmt.Sprintf("%s:%s", host, port)
}
