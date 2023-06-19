package recipe

import (
	"context"

	db "otus-recipe/app/storage/db/sqlc"

	"otus-recipe/app/api/parameters"
)

type Recipe interface {
	Get(ctx context.Context, id int64) (recipe db.Recipe, err error)
	List(ctx context.Context, limit int64, offset int64) (recipes []db.Recipe, err error)
	ListCount(ctx context.Context) (int64, error)
	Create(ctx context.Context, params *parameters.RecipeCreateParams) (db.Recipe, error)
	Update(ctx context.Context, params *parameters.RecipeUpdateParams) (db.Recipe, error)
	Delete(ctx context.Context, params parameters.RecipeDeleteParams) error
}

type recipeProcessor struct {
	store db.Store
}

func NewRecipeProcessor(store db.Store) Recipe {
	return &recipeProcessor{
		store: store,
	}
}
