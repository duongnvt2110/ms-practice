package http_handler

import (
	"context"
	"fmt"
	"log"
	"ms-practice/booking-service/pkg/container"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer(c *container.Container) {
	h := gin.Default()
	addr := resolveAddr(c)
	srv := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h,
	}
	SetRoutes(h, c.Cfg, c.Usecases)
	// http_middleware.SetMiddleware(h)
	go func() {
		log.Printf("Server is running on http://%s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http server failed: %v", err)
		}
	}()

	gracefullShutdown(srv)
	log.Println("Server exiting")
}

func gracefullShutdown(srv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
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
