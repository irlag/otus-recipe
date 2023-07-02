package models

//go:generate easyjson

const (
	RecipeEventName = "recipe.notification"
)

//easyjson:json
type RecipeUpdated struct {
	Name string `json:"name"`
}
