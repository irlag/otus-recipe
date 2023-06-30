package config

import "time"

type ElasticSettings struct {
	URLS    []string
	Timeout time.Duration `envconfig:"default=20s"`
}
