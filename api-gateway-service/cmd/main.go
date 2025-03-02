package main

import (
	"net/url"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	spew.Dump("vcv")

	// Set up multiple proxy targets for /api/v1
	apiV1Targets := []*middleware.ProxyTarget{
		{URL: mustParseURL("http://host.docker.internal:8081")},
		{URL: mustParseURL("http://host.docker.internal:8082")},
		{URL: mustParseURL("http://host.docker.internal:8083")},
	}

	apiV2Targets := []*middleware.ProxyTarget{
		{URL: mustParseURL("http://host.docker.internal:8084")},
		{URL: mustParseURL("http://host.docker.internal:8085")},
		{URL: mustParseURL("http://host.docker.internal:8086")},
	}
	// Create the /api/v1 group
	apiV1 := e.Group("/v1/hello_a")
	apiV1.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer(apiV1Targets),
		Rewrite: map[string]string{
			"^/v1/hello_a/*": "/$1", // Remove "/v1/hello_a" prefix before forwarding
		},
	}))

	// Example for /api/v2 (just for reference)
	apiV2 := e.Group("/v2/hello_b")
	apiV2.Use(middleware.ProxyWithConfig(middleware.ProxyConfig{
		Balancer: middleware.NewRoundRobinBalancer(apiV2Targets),
		Rewrite: map[string]string{
			"^/v2/hello_b/*": "/$1", // Remove "/v1/hello_a" prefix before forwarding
		},
	}))

	// Start the server
	e.Logger.Fatal(e.Start(":3000"))
}

// Helper function to parse URLs
func mustParseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err) // Fail fast if the URL is invalid
	}
	return u
}
