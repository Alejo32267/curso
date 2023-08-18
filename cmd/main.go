package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// routes for REST service (apis)
	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World 2!")
	})
	e.GET("/version", version)

	e.Logger.Fatal(e.Start(":8000"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func version(c echo.Context) error {
	return c.String(http.StatusOK, "v0.0.1")
}
