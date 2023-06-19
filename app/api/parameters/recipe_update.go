package parameters

//go:generate easyjson

import (
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//easyjson:json
type RecipeUpdateParams struct {
	RecipeID      int64  `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	CookingTime   int    `json:"cooking_time"`
	Calories      int    `json:"calories"`
	Proteins      int    `json:"proteins"`
	Fats          int    `json:"fats"`
	Carbohydrates int    `json:"carbohydrates"`
	Version       string `json:"version"`
}

func NewRecipeUpdateParamsFromRequest(request *http.Request) (*RecipeUpdateParams, error) {
	vars := mux.Vars(request)
	recipeId, err := strconv.Atoi(vars["recipe_id"])
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	recipeRequest := &RecipeUpdateParams{}
	err = recipeRequest.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}

	recipeRequest.RecipeID = int64(recipeId)

	return recipeRequest, nil
}
