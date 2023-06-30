package cmd

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	"otus-recipe/app/storage/elastic/elastic_index"
)

func (c *commandProcessor) RefreshElasticRecipeIndex(ctx context.Context, log *zap.Logger) error {
	recipeIndex := elastic_index.GetRecipe()

	if err := c.elasticsearch.CreateIndex(ctx, recipeIndex); err != nil {
		return err
	}

	start := time.Now()
	err := c.elasticsearch.BulkLoadRecipes(ctx, c.store, recipeIndex)
	log.Info(fmt.Sprintf(
		"Index '%s' data load executed in %s",
		recipeIndex.GetName(),
		time.Now().Sub(start),
	))
	if err != nil {
		return err
	}

	if err := c.elasticsearch.RefreshIndex(ctx); err != nil {
		return err
	}

	if err := c.elasticsearch.UpdateAliases(ctx, recipeIndex); err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Index '%s' updated", recipeIndex.GetName()))

	return nil
}
