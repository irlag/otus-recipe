package responses

//go:generate easyjson

import (
	"net/http"

	"otus-recipe/app/builders"
	db "otus-recipe/app/storage/db/sqlc"
)

//easyjson:json
type RecipeUpdateOkResponse struct {
	RecipeCommonOkResponse
}

func NewRecipeUpdateOkResponse(recipe db.Recipe) RecipeUpdateOkResponse {
	return RecipeUpdateOkResponse{
		RecipeCommonOkResponse: RecipeCommonOkResponse{
			ID:            recipe.ID,
			Name:          recipe.Name,
			Description:   recipe.Description.String,
			CookingTime:   int(recipe.CookingTime),
			Calories:      builders.GetIntValueFromSqlNull(recipe.Calories),
			Proteins:      builders.GetIntValueFromSqlNull(recipe.Proteins),
			Fats:          builders.GetIntValueFromSqlNull(recipe.Fats),
			Carbohydrates: builders.GetIntValueFromSqlNull(recipe.Carbohydrates),
			Version:       recipe.Version.String(),
		},
	}
}

func (r *RecipeUpdateOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := r.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
