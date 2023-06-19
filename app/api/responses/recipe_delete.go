package responses

import (
	"net/http"
)

type RecipeDeleteOkResponse struct{}

func NewRecipeDeleteOkResponse() RecipeDeleteOkResponse {
	return RecipeDeleteOkResponse{}
}

func (r *RecipeDeleteOkResponse) WriteResponse(rw http.ResponseWriter) {
	WriteJsonResponse(rw, http.StatusNoContent, []byte{})
}
