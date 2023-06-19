package responses

//go:generate easyjson

import (
	"net/http"

	db "otus-recipe/app/storage/db/sqlc"
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

func NewRecipeGetOkResponse(recipe db.Recipe) RecipeGetOkResponse {
	return RecipeGetOkResponse{
		RecipeCommonOkResponse: RecipeCommonOkResponse{
			ID:            recipe.ID,
			Name:          recipe.Name,
			Description:   recipe.Description.String,
			CookingTime:   int(recipe.CookingTime),
			Calories:      getIntValueFromSqlNull(recipe.Calories),
			Proteins:      getIntValueFromSqlNull(recipe.Proteins),
			Fats:          getIntValueFromSqlNull(recipe.Fats),
			Carbohydrates: getIntValueFromSqlNull(recipe.Carbohydrates),
			Version:       recipe.Version.String(),
		},
	}
}

func (r *RecipeGetOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := r.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
