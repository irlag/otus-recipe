package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"

	"otus-recipe/app/api/parameters"
	"otus-recipe/app/models"
	"otus-recipe/app/storage/elastic/elastic_index"
)

func (e *elastic) SearchRecipes(ctx context.Context, params *parameters.RecipeListParams, paginated models.Paginated) (recipes []*elastic_index.Recipe, err error) {
	var queryJSON bytes.Buffer
	var result map[string]interface{}

	query := map[string]interface{}{}
	var querySearch map[string]interface{}

	querySort := []map[string]interface{}{
		{
			"_score": map[string]interface{}{
				"order": "desc",
			},
		},
	}

	if params.Query != "" {
		querySearch = map[string]interface{}{
			"query_string": map[string]interface{}{
				"query":  params.Query,
				"fields": []string{"name^2", "description"},
			},
		}
	}

	if querySearch != nil {
		query = map[string]interface{}{
			"query": querySearch,
			"sort":  querySort,
		}
	}

	if err := json.NewEncoder(&queryJSON).Encode(query); err != nil {
		return nil, err
	}

	response, err := e.client.Search(
		e.client.Search.WithContext(ctx),
		e.client.Search.WithIndex(elastic_index.RecipeIndexAlias),
		e.client.Search.WithBody(&queryJSON),
		e.client.Search.WithSize(int(paginated.GetLimit())),
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

	totalHits := result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)
	paginated.SetTotal(int64(totalHits))

	for _, hit := range result["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var recipe elastic_index.Recipe
		var recipeJSON []byte

		recipeJSON, err = json.Marshal(hit.(map[string]interface{})["_source"].(map[string]interface{}))
		if err != nil {
			return nil, err
		}
		err = recipe.UnmarshalJSON(recipeJSON)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, &recipe)
	}

	return recipes, nil
}
