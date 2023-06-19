package recipe

import (
	"database/sql"
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
			switch err {
			case sql.ErrNoRows:
				response = responses.NewErrorResponse(http.StatusNotFound, appErrors.RecipeNotFoundError)
			default:
				response = responses.NewErrorResponse(http.StatusInternalServerError, err)
			}
			response.WriteErrorResponse(writer)

			return
		}

		recipeOkResponse := responses.NewRecipeGetOkResponse(recipe)
		recipeOkResponse.WriteResponse(writer)
	}
}
