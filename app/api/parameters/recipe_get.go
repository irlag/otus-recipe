package parameters

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RecipeGetParams struct {
	RecipeID int64
}

func NewRecipeGetParamsFromRequest(request *http.Request) (*RecipeGetParams, error) {
	vars := mux.Vars(request)
	recipeId, err := strconv.Atoi(vars["recipe_id"])
	if err != nil {
		return nil, err
	}

	return &RecipeGetParams{
		RecipeID: int64(recipeId),
	}, nil
}
