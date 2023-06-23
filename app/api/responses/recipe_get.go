package responses

//go:generate easyjson

import (
	"net/http"
	"strconv"

	"otus-recipe/app/storage/elastic/elastic_index"
)

type RecipeCommonOkResponse struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	CookingTime   int    `json:"cooking_time"`
	Calories      *int32 `json:"calories"`
	Proteins      *int32 `json:"proteins"`
	Fats          *int32 `json:"fats"`
	Carbohydrates *int32 `json:"carbohydrates"`
	Version       string `json:"version"`
}

//easyjson:json
type RecipeGetOkResponse struct {
	RecipeCommonOkResponse
}

func NewRecipeGetOkResponse(recipe *elastic_index.Recipe) RecipeGetOkResponse {
	recipeId, _ := strconv.Atoi(recipe.ID)

	return RecipeGetOkResponse{
		RecipeCommonOkResponse: RecipeCommonOkResponse{
			ID:            int64(recipeId),
			Name:          recipe.Name,
			Description:   recipe.Description,
			CookingTime:   int(recipe.CookingTime),
			Calories:      recipe.Calories,
			Proteins:      recipe.Proteins,
			Fats:          recipe.Fats,
			Carbohydrates: recipe.Carbohydrates,
			Version:       recipe.Version,
		},
	}
}

func (r *RecipeGetOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := r.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
