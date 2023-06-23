package elastic

import (
	"context"
	"errors"
)

func (e *elastic) RefreshIndex(ctx context.Context) error {
	res, err := e.client.Indices.Refresh(
		e.client.Indices.Refresh.WithContext(ctx),
	)
	if err != nil {
		return err
	}
	defer func() {
		res.Body.Close()
	}()

	if res.IsError() {
		return errors.New(res.String())
	}

	return nil
}
