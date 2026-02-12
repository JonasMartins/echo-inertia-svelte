// Package router
package router

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"echo-inertia.com/src/services/main/internal/bootstrap"
	"echo-inertia.com/src/services/web"
	inertia "github.com/kohkimakimoto/inertia-echo/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func NewRouter(b *bootstrap.Bootstrap) *echo.Echo {

	e := echo.New()
	setLogger(e)
	setErrorHandler(e)

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

	e.GET("/login", func(c echo.Context) error {
		return inertia.Render(c, "Login", map[string]any{})
	})

	e.POST("/login", func(c echo.Context) error {
		var req LoginRequest
		if err := c.Bind(&req); err != nil {
			return inertia.Render(c, "Login", map[string]any{
				"errors": map[string]string{
					"form": "Invalid data",
				},
			})
		}
		fmt.Println(c.Request().Header.Get("Content-Type"))

		return c.Redirect(http.StatusFound, "/")
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

func setErrorHandler(e *echo.Echo) {
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError

		var he *echo.HTTPError
		if errors.As(err, &he) {
			code = he.Code
		}

		accept := c.Request().Header.Get("Accept")

		// If client does NOT expect HTML, fallback to default (API / JSON)
		if !strings.Contains(accept, "text/html") {
			e.DefaultHTTPErrorHandler(err, c)
			return
		}

		var renderErr error

		switch code {
		case http.StatusNotFound:
			renderErr = inertia.Render(c, "404", nil)

		case http.StatusUnauthorized:
			renderErr = inertia.Render(c, "401", nil)

		case http.StatusForbidden:
			renderErr = inertia.Render(c, "403", nil)

		default:
			renderErr = inertia.Render(c, "500", nil)
		}
		if renderErr != nil {
			log.Error(renderErr)
			e.DefaultHTTPErrorHandler(err, c)
		}
	}
}
