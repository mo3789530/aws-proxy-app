package server

import (
	"aws-proxy-app/internal/pkg/proxy"
	"github.com/labstack/echo/v4"
)

func SetupProxyRoutes(e *echo.Echo) {
	config := proxy.NewConfig("./config/config_sample.yaml")
	e.Any("/*", func(c echo.Context) error {
		rproxy := proxy.NewReverseProxy(config)
		return rproxy.Handler(c)
	})
}
