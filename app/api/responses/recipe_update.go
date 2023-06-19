package responses

//go:generate easyjson

import (
	"net/http"

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

func (r *RecipeUpdateOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := r.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
