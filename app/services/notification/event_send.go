package notification

import (
	"context"

	"otus-recipe/app/models"
)

func (a *Service) EventSend(ctx context.Context, message models.Message) error {
	eventData, err := message.MarshalJSON()
	if err != nil {
		return err
	}

	err = a.client.Event(ctx, models.RecipeEventName, string(eventData))
	if err != nil {
		return err
	}

	return nil
}
