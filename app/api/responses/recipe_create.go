package responses

//go:generate easyjson

import (
	"database/sql"
	"net/http"

	db "otus-recipe/app/storage/db/sqlc"
)

//easyjson:json
type RecipeCreateOkResponse struct {
	RecipeCommonOkResponse
}

func NewRecipeCreateOkResponse(recipe db.Recipe) RecipeCreateOkResponse {
	return RecipeCreateOkResponse{
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

func getIntValueFromSqlNull(prop sql.NullInt32) *int32 {
	if prop.Valid {
		return &prop.Int32
	}

	return nil
}

func (r *RecipeCreateOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := r.MarshalJSON()

	WriteJsonResponse(rw, http.StatusCreated, payload)
}
