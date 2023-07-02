package processors

import (
	"go.uber.org/zap"

	"otus-recipe/app/processors/cmd"
	"otus-recipe/app/services"
	"otus-recipe/app/storage/elastic"

	"otus-recipe/app/config"
	"otus-recipe/app/processors/recipe"
	db "otus-recipe/app/storage/db/sqlc"
)

type Processors struct {
	HealthcheckProcessor HealthcheckProcessor
	RecipeProcessor      recipe.Recipe
	Command              cmd.Command
}

func NewProcessor(
	store db.Store,
	services *services.Services,
	elasticsearch elastic.Elastic,
	log *zap.Logger,
	config *config.Config,
) *Processors {
	return &Processors{
		HealthcheckProcessor: NewHealtcheckProcessor(),
		RecipeProcessor:      recipe.NewRecipeProcessor(store, elasticsearch, services),
		Command:              cmd.New(store, elasticsearch),
	}
}
