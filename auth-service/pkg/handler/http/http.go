package http_handler

import (
	"context"
	"os"
	"os/signal"
	"time"

	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/container"
	"ms-practice/auth-service/pkg/handler/http/auth"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func StartHTTPServer(c *container.Container) {
	e := echo.New()
	setRoutes(e, c)
	initialMiddleware(e)
	runHttpServer(e, c.Cfg)
}

func setRoutes(e *echo.Echo, c *container.Container) {
	// Version
	apiV1 := e.Group("/v1")

	// Auth Handler
	authHandler := auth.NewAuthHandler(c.Cfg, c.Validate, c.Usecase)

	// Token Routes
	apiV1.POST("/register", authHandler.Register)
	apiV1.POST("/login", authHandler.Login)
	apiV1.POST("/logout", authHandler.Logout)

	// SSO Routes
	gAuth := apiV1.Group("/auth/google")
	gAuth.GET("/login", authHandler.OauthGoogleLogin)
	gAuth.GET("/callback", authHandler.OauthGoogleCallback).Name = "oauth.callback"
}

func initialMiddleware(e *echo.Echo) {
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
}

func authMiddleware(e *echo.Echo) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		// ...
		SigningKey: []byte("secret"),
		// ...
	})
}

func runHttpServer(e *echo.Echo, cfg *config.Config) {
	//Listen graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start http server
	errCh := make(chan error)
	go func() {
		errCh <- e.Start(":" + cfg.App.Port)
	}()

	// Handle channel Err by Server and Graceful Shutdown
	select {
	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := e.Shutdown(ctx)
		if err != nil {
			log.Fatal("Server shutdown failed:", err)
		}
		log.Info("Server exited gracefully")
	case err := <-errCh:
		log.Info("Server shutdown by", err)
	}

}
