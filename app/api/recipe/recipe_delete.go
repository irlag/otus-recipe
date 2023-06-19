package recipe

import (
	"database/sql"
	"net/http"

	"otus-recipe/app/api/parameters"
	"otus-recipe/app/api/responses"
	appErrors "otus-recipe/app/errors"
)

func (r *Recipe) Delete() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		recipeDeleteParams, err := parameters.NewRecipeDeleteParamsFromRequest(request)
		if err != nil {
			responses.NewErrorResponse(http.StatusBadRequest, err).WriteErrorResponse(writer)

			return
		}

		_, err = r.processors.RecipeProcessor.Get(request.Context(), recipeDeleteParams.RecipeID)
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

		err = r.processors.RecipeProcessor.Delete(request.Context(), recipeDeleteParams)
		if err != nil {
			responses.NewErrorResponse(http.StatusInternalServerError, err).WriteErrorResponse(writer)

			return
		}

		recipeDeleteOkResponse := responses.NewRecipeDeleteOkResponse()
		recipeDeleteOkResponse.WriteResponse(writer)
	}
}
