package elastic

import (
	"context"
	"errors"
	"strings"

	"otus-recipe/app/storage/elastic/elastic_index"
)

func (e *elastic) CreateIndex(ctx context.Context, index elastic_index.Index) error {
	mapping := elastic_index.Mapping{}

	err := mapping.UnmarshalJSON([]byte(index.GetSettings()))
	if err != nil {
		return err
	}

	err = mapping.UnmarshalJSON([]byte(index.GetMappings()))
	if err != nil {
		return err
	}

	indexScheme, err := mapping.MarshalJSON()
	if err != nil {
		return err
	}

	res, err := e.client.Indices.Create(
		index.GetName(),
		e.client.Indices.Create.WithBody(strings.NewReader(string(indexScheme))),
		e.client.Indices.Create.WithContext(ctx),
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
