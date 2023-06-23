package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"otus-recipe/app/storage/elastic/elastic_index"
)

func (e *elastic) UpdateAliases(ctx context.Context, index elastic_index.Index) error {
	query := map[string]interface{}{}
	oldIndex, err := e.resolveIndex(ctx, index)
	if err != nil {
		return err
	}
	if oldIndex == nil {
		query = map[string]interface{}{
			"actions": map[string]interface{}{
				"add": map[string]interface{}{
					"index": index.GetName(),
					"alias": index.GetAlias(),
				},
			},
		}
	} else {
		query = map[string]interface{}{
			"actions": []map[string]interface{}{
				{
					"add": map[string]interface{}{
						"index": index.GetName(),
						"alias": index.GetAlias(),
					},
				},
				{
					"remove_index": map[string]interface{}{
						"index": oldIndex,
					},
				},
			},
		}
	}

	var buf bytes.Buffer
	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		return err
	}

	res, err := e.client.Indices.UpdateAliases(
		&buf,
		e.client.Indices.UpdateAliases.WithContext(ctx),
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

func (e *elastic) resolveIndex(ctx context.Context, index elastic_index.Index) (*interface{}, error) {
	res, err := e.client.Indices.ResolveIndex(
		[]string{index.GetAlias()},
		e.client.Indices.ResolveIndex.WithContext(ctx),
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		res.Body.Close()
	}()

	if res.IsError() {
		return nil, errors.New(res.String())
	}

	var r map[string][]map[string]interface{}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return nil, err
	}

	aliases := r["aliases"]
	if len(aliases) == 0 {
		return nil, nil
	}

	return &aliases[0]["indices"].([]interface{})[0], nil
}
