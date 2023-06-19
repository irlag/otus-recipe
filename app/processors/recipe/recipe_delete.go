package recipe

import (
	"context"

	appErrors "otus-recipe/app/errors"

	"otus-recipe/app/api/parameters"
)

func (r *recipeProcessor) Delete(ctx context.Context, params parameters.RecipeDeleteParams) error {
	rowsAffected, err := r.store.DeleteRecipe(ctx, params.RecipeID)
	if err != nil {
		return err
	}

	if rowsAffected <= 0 {
		return appErrors.RecipeNotDeletedError
	}

	return nil
}
