package utils

import (
	"AceofHearts/config"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var (
	httpClient *http.Client
	clientOnce sync.Once
)

// GetHTTPClient 返回全局复用的 HTTP 客户端（连接池 + 代理）
func GetHTTPClient() *http.Client {
	clientOnce.Do(func() {
		transport := &http.Transport{
			MaxIdleConns:        20,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		}

		if config.ProxyURL != "" {
			proxyURL, err := url.Parse(config.ProxyURL)
			if err == nil {
				transport.Proxy = http.ProxyURL(proxyURL)
			}
		}

		httpClient = &http.Client{
			Transport: transport,
			Timeout:   30 * time.Minute,
		}
	})
	return httpClient
}

// ResetHTTPClient 代理变更后重置客户端（在 Setting 阶段调用）
func ResetHTTPClient() {
	clientOnce = sync.Once{}
	httpClient = nil
}
