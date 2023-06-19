package api

import "errors"

var (
	RouteNotExistError = errors.New("route not found")
)

type Route struct {
	Name   string
	Method string
	Path   string
	Secure bool
}

type Routes map[string]Route

func (r *Routes) Find(name string) (Route, error) {
	if appRoute, ok := AppRoutes[name]; ok {
		return appRoute, nil
	}
	return Route{}, RouteNotExistError
}

var AppRoutes = Routes{
	"metrics": Route{
		Name:   "metrics",
		Method: "GET",
		Path:   "/metrics",
		Secure: false,
	},

	"healthcheck": Route{
		Name:   "healthcheck",
		Method: "GET",
		Path:   "/health",
		Secure: false,
	},

	"recipe_create": Route{
		Name:   "recipe_create",
		Method: "POST",
		Path:   "/recipe",
		Secure: false,
	},
	"recipe_list": Route{
		Name:   "recipe_list",
		Method: "GET",
		Path:   "/recipe",
		Secure: false,
	},
	"recipe_detail": Route{
		Name:   "recipe_detail",
		Method: "GET",
		Path:   "/recipe/{recipe_id:[0-9]+}",
		Secure: false,
	},
	"recipe_update": Route{
		Name:   "recipe_update",
		Method: "PUT",
		Path:   "/recipe/{recipe_id:[0-9]+}",
		Secure: false,
	},
	"recipe_delete": Route{
		Name:   "recipe_delete",
		Method: "DELETE",
		Path:   "/recipe/{recipe_id:[0-9]+}",
		Secure: false,
	},
}
