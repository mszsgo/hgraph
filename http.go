package hgraph

import (
	"net/http"
	"net/url"
	"os"
	"time"
)

func HttpClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			// 服务HTTP代理配置，示例系统环境变量： "MS_HTTP_PROXY=211.152.57.29:39084"
			Proxy: func(request *http.Request) (url *url.URL, err error) {
				msHttpProxy := os.Getenv("MS_HTTP_PROXY")
				if msHttpProxy != "" {
					request.URL.Host = msHttpProxy
				}
				return request.URL, err
			},
		},
		Timeout: 5 * time.Second,
	}
}
