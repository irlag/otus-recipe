package recipe

import (
	"github.com/gorilla/mux"

	"otus-recipe/app/api"
	"otus-recipe/app/processors"
)

type Recipe struct {
	processors *processors.Processors
}

func NewRecipeApi(processors *processors.Processors) *Recipe {
	return &Recipe{
		processors: processors,
	}
}

func (r *Recipe) HandleMethods(router *mux.Router) {
	router.HandleFunc(api.AppRoutes["recipe_list"].Path, r.List()).
		Methods(api.AppRoutes["recipe_list"].Method).
		Name(api.AppRoutes["recipe_list"].Name)

	router.HandleFunc(api.AppRoutes["recipe_detail"].Path, r.Get()).
		Methods(api.AppRoutes["recipe_detail"].Method).
		Name(api.AppRoutes["recipe_detail"].Name)

	router.HandleFunc(api.AppRoutes["recipe_create"].Path, r.Create()).
		Methods(api.AppRoutes["recipe_create"].Method).
		Name(api.AppRoutes["recipe_create"].Name)

	router.HandleFunc(api.AppRoutes["recipe_delete"].Path, r.Delete()).
		Methods(api.AppRoutes["recipe_delete"].Method).
		Name(api.AppRoutes["recipe_delete"].Name)

	router.HandleFunc(api.AppRoutes["recipe_update"].Path, r.Update()).
		Methods(api.AppRoutes["recipe_update"].Method).
		Name(api.AppRoutes["recipe_update"].Name)
}
