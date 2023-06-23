package recipe

import (
	"context"

	"otus-recipe/app/api/parameters"
	"otus-recipe/app/models"
	"otus-recipe/app/storage/elastic/elastic_index"
)

func (r *recipeProcessor) List(ctx context.Context, paginated models.Paginated, params *parameters.RecipeListParams) (recipes []*elastic_index.Recipe, err error) {
	recipes, err = r.elasticsearch.SearchRecipes(ctx, params, paginated)
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
