package parameters

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type RecipeDeleteParams struct {
	RecipeID int64
}

func NewRecipeDeleteParamsFromRequest(request *http.Request) (RecipeDeleteParams, error) {
	vars := mux.Vars(request)
	recipeId, err := strconv.Atoi(vars["recipe_id"])
	if err != nil {
		return RecipeDeleteParams{}, err
	}

	return RecipeDeleteParams{
		RecipeID: int64(recipeId),
	}, nil
}
