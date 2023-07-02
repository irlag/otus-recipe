package clients

import (
	"bytes"
	"context"
	"net/http"

	"otus-recipe/app/config"
	"otus-recipe/app/models"
)

type Notification interface {
	Event(ctx context.Context, eventName string, data string) error
}

type notification struct {
	client *http.Client
	host   string
}

func (n notification) Event(ctx context.Context, eventName string, data string) error {
	event := models.Event{
		Name: eventName,
		Data: data,
	}

	body, err := event.MarshalJSON()
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(body)

	r, err := http.NewRequest(http.MethodPost, n.host+"/event", bodyReader)
	if err != nil {
		return err
	}

	r = r.WithContext(ctx)

	resp, err := n.client.Do(r)
	if err != nil && resp == nil {
		return err
	}

	return nil
}

func newNotificationHttpClient(service *config.Service) Notification {
	return &notification{
		client: newHttpClient(&config.Service{
			Host:           service.Host,
			RequestTimeout: service.RequestTimeout,
		}),
		host: service.Host,
	}
}
