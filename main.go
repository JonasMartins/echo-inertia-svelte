package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Health check endpoint
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Greeting endpoint
	e.GET("/greeting", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello! Your Echo server is up and running.")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
