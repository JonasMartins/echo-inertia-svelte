// Package router
package router

import (
	"echo-inertia.com/src/services/main/internal/bootstrap"
	"github.com/labstack/echo/v4"
)

func NewRouter(b *bootstrap.Bootstrap) *echo.Echo {

	e := echo.New()

	return e
}
