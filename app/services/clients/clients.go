package clients

import (
	"net"
	"net/http"
	"time"

	"otus-recipe/app/config"
)

type Clients struct {
	Notification Notification
}

func New(config *config.Config) *Clients {
	return &Clients{
		Notification: newNotificationHttpClient(config.Clients.Notification),
	}
}

func newHttpClient(service *config.Service) *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   service.RequestTimeout,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   1 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			ResponseHeaderTimeout: service.RequestTimeout,
		},
		Timeout: service.RequestTimeout,
	}
}
