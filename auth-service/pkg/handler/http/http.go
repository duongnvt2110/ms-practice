package http_handler

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"time"

	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/container"
	"ms-practice/auth-service/pkg/handler/http/auth"
	"ms-practice/auth-service/pkg/usecase"
	"ms-practice/auth-service/pkg/utils/apperr"

	resp "ms-practice/pkg/http/echo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func StartHTTPServer(ctx context.Context, c *container.Container) {
	e := echo.New()
	registerRoutes(e, c)
	registerMiddleware(e, c)
	startHttpServer(e, c.Cfg)
}

func registerRoutes(e *echo.Echo, c *container.Container) {
	// Version
	apiV1 := e.Group("/v1")

	// Auth Handler
	authHandler := auth.NewAuthHandler(c.Cfg, c.Validate, c.Usecase)

	// Inital Middlewares
	authMD := authMiddleware(c.Usecase.AuthProfileUC)

	// Token Routes
	authRoutes := apiV1.Group("/auths")
	authRoutes.POST("/register", authHandler.Register)
	authRoutes.POST("/login", authHandler.Login)
	authRoutes.POST("/refresh_token", authHandler.RefreshToken)
	authRoutes.POST("/logout", authHandler.Logout, authMD)

	// SSO Routes
	gAuth := authRoutes.Group("/google")
	gAuth.GET("/login", authHandler.OauthGoogleLogin)
	gAuth.GET("/callback", authHandler.OauthGoogleCallback).Name = "oauth.callback"
}

func registerMiddleware(e *echo.Echo, c *container.Container) {
	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
}

func authMiddleware(authUC usecase.AuthProfileUC) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Get Token
			auth := c.Request().Header.Get(echo.HeaderAuthorization)
			token := strings.TrimPrefix(auth, "Bearer")
			token = strings.TrimSpace(token)
			// Validate Token
			authClaims, err := authUC.ValidateToken(token)
			if err != nil {
				return resp.ResponseWithError(c, apperr.ErrInvalidToken)
			}
			c.Set("auth_profile_id", authClaims.AuthProfileId)
			return next(c)
		}
	}
}

func startHttpServer(e *echo.Echo, cfg *config.Config) {
	//Listen graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Start http server
	errCh := make(chan error)
	go func() {
		defer close(errCh)
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
		return
	case err := <-errCh:
		log.Info("Server shutdown by", err)
		return
	}

}
