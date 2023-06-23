package cmd

import (
	"context"

	"go.uber.org/zap"

	db "otus-recipe/app/storage/db/sqlc"
	"otus-recipe/app/storage/elastic"
)

type Command interface {
	RefreshElasticRecipeIndex(ctx context.Context, log *zap.Logger) error
}

type commandProcessor struct {
	elasticsearch elastic.Elastic
	store         db.Store
}

func New(
	store db.Store,
	elasticsearch elastic.Elastic,
) Command {
	return &commandProcessor{
		elasticsearch: elasticsearch,
		store:         store,
	}
}
