package recipe

import (
	"context"

	"otus-recipe/app/storage/elastic/elastic_index"
)

func (r *recipeProcessor) Get(ctx context.Context, id int64) (recipe *elastic_index.Recipe, err error) {
	recipe, err = r.elasticsearch.GetRecipe(ctx, id)
	if err != nil {
		return recipe, err
	}

	return recipe, nil
}
