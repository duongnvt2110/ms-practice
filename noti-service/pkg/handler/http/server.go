package http_handler

import (
	"context"
	"fmt"
	"log"
	"ms-practice/noti-service/pkg/container"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer(ctx context.Context, c *container.Container) {
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	addr := fmt.Sprintf("%s:%s", c.Cfg.App.Host, c.Cfg.App.Port)
	srv := &http.Server{Addr: addr, Handler: router}

	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = srv.Shutdown(shutdownCtx)
	}()

	log.Printf("noti service http listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("http server error: %v", err)
	}
}
