package services

import (
	"otus-recipe/app/services/clients"
	"otus-recipe/app/services/notification"
)

type Services struct {
	Notification notification.Notification
}

func New(clients *clients.Clients) *Services {
	return &Services{
		Notification: notification.New(
			notification.Config{
				Client: clients.Notification,
			},
		),
	}
}
