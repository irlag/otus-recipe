package parameters

//go:generate easyjson

import (
	"net/http"
	"strconv"
)

//easyjson:json
type RecipeListParams struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

func NewRecipeListParamsFromRequest(request *http.Request) *RecipeListParams {
	queryParams := request.URL.Query()

	query := queryParams.Get("query")
	limit, _ := strconv.Atoi(queryParams.Get("limit"))
	page, _ := strconv.Atoi(queryParams.Get("page"))

	return &RecipeListParams{
		Query: query,
		Limit: limit,
		Page:  page,
	}

}
