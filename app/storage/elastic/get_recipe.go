package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"otus-recipe/app/storage/elastic/elastic_index"
)

func (e *elastic) GetRecipe(ctx context.Context, recipeId int64) (*elastic_index.Recipe, error) {
	var queryJSON bytes.Buffer
	var result map[string]interface{}

	query := map[string]interface{}{}
	var querySearch map[string]interface{}

	querySearch = map[string]interface{}{
		"match": map[string]interface{}{
			"id": recipeId,
		},
	}

	query = map[string]interface{}{
		"query": querySearch,
	}

	if err := json.NewEncoder(&queryJSON).Encode(query); err != nil {
		return nil, err
	}

	response, err := e.client.Search(
		e.client.Search.WithContext(ctx),
		e.client.Search.WithIndex(elastic_index.RecipeIndexAlias),
		e.client.Search.WithBody(&queryJSON),
	)
	if err != nil {
		return nil, err
	}
	defer func() {
		response.Body.Close()
	}()

	if response.IsError() {
		return nil, errors.New(response.String())
	}

	if err = json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, err
	}

	var recipe elastic_index.Recipe

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var recipeJSON []byte

		recipeJSON, err = json.Marshal(hit.(map[string]interface{})["_source"].(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		err = recipe.UnmarshalJSON(recipeJSON)
		if err != nil {
			return nil, err
		}
	}

	return &recipe, nil
}
