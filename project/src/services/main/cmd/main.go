// Package main
package main

import (
	inertia "github.com/kohkimakimoto/inertia-echo/v2"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// 1. Initialize the built-in HTML Renderer
	r := inertia.NewHTMLRenderer()

	// 4. Set up Middleware
	e.Use(inertia.MiddlewareWithConfig(inertia.MiddlewareConfig{
		Renderer: r,
	}))

	e.GET("/", func(c echo.Context) error {
		return inertia.Render(c, "Home", map[string]any{
			"message": "Welcome to the Home Page!",
		})
	})

	e.GET("/about", func(c echo.Context) error {
		return inertia.Render(c, "About", map[string]any{
			"content": "We are building a highly performant Go web app using Svelte 5.",
		})
	})

	e.Logger.Fatal(e.Start(":8080"))
}
