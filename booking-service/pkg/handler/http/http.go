package http_handler

import (
	"booking-service/pkg/container"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func StartHTTPServer(c *container.Container) {
	h := gin.Default()
	srv := &http.Server{
		Addr:         ":3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h,
	}
	SetRoutes(h, c.Cfg, c.Kafka)
	// http_middleware.SetMiddleware(h)
	go func() {
		log.Printf("Server is running on http://%s", c.Cfg.App.Host)
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
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
