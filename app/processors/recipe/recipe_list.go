package recipe

import (
	"context"

	db "otus-recipe/app/storage/db/sqlc"
)

func (r *recipeProcessor) List(ctx context.Context, limit int64, offset int64) (recipes []db.Recipe, err error) {
	listDbRecipeParams := db.ListRecipeParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	recipes, err = r.store.ListRecipe(ctx, listDbRecipeParams)
	if err != nil {
		return recipes, err
	}

	return recipes, nil
}

func (r *recipeProcessor) ListCount(ctx context.Context) (int64, error) {
	count, err := r.store.ListRecipeCount(ctx)
	if err != nil {
		return 0, err
	}

	return count, nil
}
