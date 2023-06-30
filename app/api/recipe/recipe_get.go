package recipe

import (
	"net/http"

	"otus-recipe/app/api/parameters"
	"otus-recipe/app/api/responses"
	appErrors "otus-recipe/app/errors"
)

func (r *Recipe) Get() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		recipeGetParams, err := parameters.NewRecipeGetParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		recipe, err := r.processors.RecipeProcessor.Get(request.Context(), recipeGetParams.RecipeID)
		if err != nil {
			response := &responses.ErrorResponse{}
			response = responses.NewErrorResponse(http.StatusInternalServerError, err)
			response.WriteErrorResponse(writer)

			return
		}
		if recipe == nil {
			response := &responses.ErrorResponse{}
			response = responses.NewErrorResponse(http.StatusNotFound, appErrors.RecipeNotFoundError)
			response.WriteErrorResponse(writer)

			return
		}

		recipeOkResponse := responses.NewRecipeGetOkResponse(recipe)
		recipeOkResponse.WriteResponse(writer)
	}
}
