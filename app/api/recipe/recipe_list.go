package recipe

import (
	"net/http"

	"otus-recipe/app/api/parameters"
	"otus-recipe/app/api/responses"
	"otus-recipe/app/models"
)

func (r *Recipe) List() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		recipeListParams := parameters.NewRecipeListParamsFromRequest(request)
		paginated := models.NewPaginatedFromRequest(recipeListParams)

		recipes, err := r.processors.RecipeProcessor.List(request.Context(), paginated, recipeListParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		recipeListResponse := responses.NewRecipeListOkResponse(recipes, paginated)
		recipeListResponse.WriteResponse(writer)
	}
}
