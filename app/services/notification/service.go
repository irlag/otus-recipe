package notification

import (
	"context"

	"otus-recipe/app/models"
	"otus-recipe/app/services/clients"
)

type Config struct {
	Client clients.Notification
}

type Notification interface {
	EventSend(ctx context.Context, message models.Message) error
}

type Service struct {
	client clients.Notification
}

func New(cfg Config) Notification {
	return &Service{
		client: cfg.Client,
	}
}
