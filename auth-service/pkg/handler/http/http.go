package http_handler

import (
	"auth-service/pkg/config"
	"auth-service/pkg/container"
	"auth-service/pkg/handler/http/auth"
	"context"
	"os"
	"os/signal"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func StartHTTPServer(c *container.Container) {
	e := echo.New()
	setRoutes(e, c.Cfg)
	initialMiddleware(e)
	runHttpServer(e, c.Cfg)
}

func setRoutes(e *echo.Echo, cfg *config.Config) {
	// Version
	apiV1 := e.Group("/v1")

	//

	// Auth Handler
	authHandler := auth.NewAuthHandler(cfg)

	// Token Routes
	tAuth := apiV1.Group("token")
	tAuth.POST("/register", authHandler.Register)
	tAuth.POST("/login", authHandler.Login)
	tAuth.POST("/logout", authHandler.Logout)

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
