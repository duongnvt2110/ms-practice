package http_handler

import (
	"auth-service/pkg/config"
	"auth-service/pkg/container"
	"auth-service/pkg/handler/http/auth"
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

func StartHTTPServer(c *container.Container) {
	e := echo.New()
	// s := http.Server{
	// 	Addr:    ":3000",
	// 	Handler: e,
	// }
	// errs := make(chan error)
	setRoutes(e, c.Cfg)
	// setMiddleware(e)
	// gracefullShutdown(e)
	// go func() {
	e.Logger.Fatal(e.Start(":3000"))
	// e.Logger.Info("Server is running on http://%s", c.Cfg.App.Host)
	// err := e.ListenAndServe()
	// errs <- err
	// if err != nil && err != http.ErrServerClosed {
	// 	e.Logger.Fatal("shutting down the server")
	// }
	// }()
	// e.Logger.Info("exit", <-errs)

}

func setRoutes(e *echo.Echo, cfg *config.Config) {
	authHandler := auth.NewAuthHandler(cfg)
	e.GET("/auth/google/login", authHandler.OauthGoogleLogin)
	e.GET("/auth/google/callback", authHandler.OauthGoogleCallback).Name = "oauth.calback"
}

func setMiddleware(e *echo.Echo) {
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
}

func gracefullShutdown(e *echo.Echo) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := e.Shutdown(ctx)
	if err != nil {
		e.Logger.Fatal("Server shutdown failed:", err)
	}
	e.Logger.Info("Server exited gracefully")
}
