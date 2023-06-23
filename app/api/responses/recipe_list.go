package responses

//go:generate easyjson

import (
	"net/http"

	"otus-recipe/app/models"
	"otus-recipe/app/storage/elastic/elastic_index"
)

//easyjson:json
type RecipeListOkResponse struct {
	Paginated
	Items []RecipeGetOkResponse `json:"items"`
}

func NewRecipeListOkResponse(recipes []*elastic_index.Recipe, paginated models.Paginated) RecipeListOkResponse {
	response := RecipeListOkResponse{
		Paginated: Paginated{
			Page:  paginated.GetPage(),
			Limit: paginated.GetLimit(),
			Total: paginated.GetTotal(),
			Pages: paginated.GetPages(),
		},
		Items: []RecipeGetOkResponse{},
	}
	for _, recipe := range recipes {
		response.Items = append(response.Items, NewRecipeGetOkResponse(recipe))
	}

	return response
}

func (r *RecipeListOkResponse) WriteResponse(rw http.ResponseWriter) {
	payload, _ := r.MarshalJSON()

	WriteJsonResponse(rw, http.StatusOK, payload)
}
