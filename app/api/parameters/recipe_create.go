package parameters

//go:generate easyjson

import (
	"io"
	"net/http"
)

//easyjson:json
type RecipeCreateParams struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	CookingTime   int    `json:"cooking_time"`
	Calories      int    `json:"calories"`
	Proteins      int    `json:"proteins"`
	Fats          int    `json:"fats"`
	Carbohydrates int    `json:"carbohydrates"`
}

func NewRecipeCreateParamsFromRequest(request *http.Request) (*RecipeCreateParams, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	recipeRequest := &RecipeCreateParams{}
	err = recipeRequest.UnmarshalJSON(body)
	if err != nil {
		return nil, err
	}

	return recipeRequest, nil
}
