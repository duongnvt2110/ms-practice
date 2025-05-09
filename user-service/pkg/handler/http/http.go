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

func StartHTTPServer(c *container.Container, ctx context.Context) {
	h := mux.NewRouter()
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", c.Cfg.App.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h,
	}
	// errs := make(chan error)
	SetRoutes(h, c)
	// http_middleware.SetMiddleware(h)
	logChannel := make(chan string)

	go func() {
		logChannel <- fmt.Sprintf("Server is running on http://%s:%s", c.Cfg.App.Host, c.Cfg.App.Port)
		err := srv.ListenAndServe()
		if err != nil {
			logChannel <- err.Error()
		}
	}()

	go func() {
		for msg := range logChannel {
			fmt.Println(msg)
		}
		close(logChannel)
	}()

	// Gracefull shutdown
	<-ctx.Done()
	gracefullShutdown(srv)
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
