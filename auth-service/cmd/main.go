package main

import (
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
)

func main() {
	spew.Dump("2323232323")
	e := echo.New()
	// c := container.InitializeContainer()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World 12342341!")
	})
	errs := make(chan error)
	e.Logger.SetOutput(os.Stdout)
	go func() {
		errs <- e.Start(":3000")
	}()
	e.Logger.Info("exit",<-errs)
	// Run HTTP Server
	// http_handler.StartHTTPServer(c)
	// select {}
}

func new() {
	spew.Dump("e := echo.New()sdafsdf")
}
