package http_handler

import (
	"context"
	"fmt"
	"ms-practice/catalog-service/pkg/container"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer(c *container.Container, ctx context.Context) {
	engine := gin.Default()
	SetRoutes(engine, c)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", c.Cfg.App.Port),
		Handler:      engine,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	<-ctx.Done()
	shutdownContext, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = srv.Shutdown(shutdownContext)
}
