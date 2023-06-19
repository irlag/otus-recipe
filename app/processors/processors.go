package processors

import (
	"go.uber.org/zap"

	"otus-recipe/app/services"

	"otus-recipe/app/config"
	"otus-recipe/app/processors/recipe"
	db "otus-recipe/app/storage/db/sqlc"
)

type Processors struct {
	HealthcheckProcessor HealthcheckProcessor
	RecipeProcessor      recipe.Recipe
}

func NewProcessor(
	store db.Store,
	services *services.Services,
	log *zap.Logger,
	config *config.Config,
) *Processors {
	return &Processors{
		HealthcheckProcessor: NewHealtcheckProcessor(),
		RecipeProcessor:      recipe.NewRecipeProcessor(store),
	}
}
