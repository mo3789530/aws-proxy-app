package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupRootRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"health": "ok",
		})
	})
}
