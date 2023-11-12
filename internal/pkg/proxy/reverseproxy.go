package proxy

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
)

type ReverseProxy struct {
	Config Root
}

func NewReverseProxy(config Root) *ReverseProxy {
	return &ReverseProxy{
		Config: config,
	}
}

// ServeHttpWithPort TODO
// Proxy mode
func (r *ReverseProxy) ServeHttpWithPort() {

}

// S3 mode
func (r *ReverseProxy) mockServerWithS3(c echo.Context, bucket Bucket) error {
	return c.JSON(http.StatusOK, map[string]string{
		"bucket": bucket.BucketName,
	})
}

func (r *ReverseProxy) findConfig(sid string) *SID {
	for _, v := range r.Config.Keys {
		if v.SID.Name == sid {
			return &v.SID
		}
	}
	return nil
}

func (r *ReverseProxy) Handler(c echo.Context) error {
	host := c.Request().Host
	sid := c.Request().Header.Get("SID")
	slog.Info(fmt.Sprintf("ReverseProxy Handler %v %v", host, sid))
	cnf := r.findConfig(sid)
	if cnf == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "config not found",
		})
	}
	if cnf.Config.Storage.StorageMode == true {
		return r.mockServerWithS3(c, cnf.Config.Storage.Bucket)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"health": "ok",
	})
}
