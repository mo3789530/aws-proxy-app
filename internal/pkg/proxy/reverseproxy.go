package proxy

import (
	"fmt"
	"log/slog"
	"net/http"

	client "aws-proxy-app/internal/pkg/client"

	"github.com/labstack/echo/v4"
)

type ReverseProxy struct {
	Config  Root
	Storage client.StorageClient
}

func NewReverseProxy(config Root, storage client.StorageClient) *ReverseProxy {
	return &ReverseProxy{
		Config:  config,
		Storage: storage,
	}
}

// ServeHttpWithPort TODO
// Proxy mode
func (r *ReverseProxy) ServeHttpWithPort() error {
	return fmt.Errorf("NotImplementedError")
}

// S3 mode
func (r *ReverseProxy) mockServerWithS3(c echo.Context, bucket Bucket, sid string) error {
	// TODO
	// method := c.Request().Method
	// fmt.Println(method)
	url := c.Request().URL
	fmt.Println(url)
	fmt.Println(url.RawPath)
	bucketList, _ := r.Storage.ListBuckets()
	for _, v := range bucketList {
		fmt.Println(v)
	}

	r.Storage.GetObject(bucket.BucketName, getS3Path(sid, url.String()))

	return c.JSON(http.StatusAccepted, map[string]string{})

}

func getS3Path(sid string, url string) string {
	return fmt.Sprintf("%v/%v", sid, url)
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
	if cnf.Config.Storage.StorageMode {
		return r.mockServerWithS3(c, cnf.Config.Storage.Bucket, sid)
	}
	return c.JSON(http.StatusOK, map[string]string{
		"health": "ok",
	})
}
