package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupProxyRoutes(e *echo.Echo) {
	e.Any("/*", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{})
	})
}
