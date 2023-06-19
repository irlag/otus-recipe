package recipe

import (
	"net/http"

	"otus-recipe/app/api/responses"
	"otus-recipe/app/models"
)

func (r *Recipe) List() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		paginated := models.NewPaginatedFromRequest(request)

		recipes, err := r.processors.RecipeProcessor.List(request.Context(), paginated.GetLimit(), paginated.GetOffset())
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}
		total, err := r.processors.RecipeProcessor.ListCount(request.Context())
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		paginated.SetTotal(total)

		recipeListResponse := responses.NewRecipeListOkResponse(recipes, paginated)
		recipeListResponse.WriteResponse(writer)
	}
}
