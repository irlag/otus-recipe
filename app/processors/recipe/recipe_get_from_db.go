package recipe

import (
	"context"

	db "otus-recipe/app/storage/db/sqlc"
)

func (r *recipeProcessor) GetFromDb(ctx context.Context, id int64) (recipe db.Recipe, err error) {
	recipe, err = r.store.GetRecipe(ctx, id)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}
