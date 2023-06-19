package recipe

import (
	"net/http"

	"otus-recipe/app/api/parameters"
	"otus-recipe/app/api/responses"
)

func (r *Recipe) Create() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		recipeCreateParams, err := parameters.NewRecipeCreateParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		recipe, err := r.processors.RecipeProcessor.Create(request.Context(), recipeCreateParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		recipeOkResponse := responses.NewRecipeCreateOkResponse(recipe)
		recipeOkResponse.WriteResponse(writer)
	}
}
