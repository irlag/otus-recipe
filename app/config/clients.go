package config

import "time"

type Service struct {
	Host           string
	RequestTimeout time.Duration `envconfig:"default=1s"`
}
