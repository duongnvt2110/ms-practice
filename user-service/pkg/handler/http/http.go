package http_handler

import (
	"context"
	"log"
	"ms-practice/user-service/pkg/container"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

func StartHTTPServer(c *container.Container) {
	h := mux.NewRouter()
	srv := &http.Server{
		Addr:         ":3000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      h,
	}
	// errs := make(chan error)
	SetRoutes(h, c.Cfg)
	// http_middleware.SetMiddleware(h)
	// ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	// defer stop()
	go func() {
		log.Printf("Server is running on http://%s", c.Cfg.App.Host)
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// go func() {
	// 	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	// 	<-ctx.Done()
	// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 	defer cancel()
	// 	if err := srv.Shutdown(ctx); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()
	// gracefullShutdown(srv)
	// log.Fatal("exit", <-errs)
}

func gracefullShutdown(srv *http.Server) {
	// Implement graceful shutdown.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		log.Println("Close another connection")
		cancel()
	}()
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server shutdown faile23d tes:", err)
	}
	log.Println("Server exited gracefully")
}
