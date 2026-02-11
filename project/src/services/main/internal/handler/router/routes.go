// Package router
package router

import (
	"fmt"

	"echo-inertia.com/src/services/main/internal/bootstrap"
	"echo-inertia.com/src/services/web"
	inertia "github.com/kohkimakimoto/inertia-echo/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(b *bootstrap.Bootstrap) *echo.Echo {

	e := echo.New()
	setLogger(e)

	// 1. Initialize the built-in HTML Renderer
	r := inertia.NewHTMLRenderer()

	web.SetRender(e, r, b.Config)

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
	return e
}

func setLogger(e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogMethod: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("%s: %s -> :%d\n", v.Method, v.URI, v.Status)
			return nil
		},
	}))
}
