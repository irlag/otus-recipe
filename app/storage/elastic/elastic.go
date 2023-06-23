package elastic

import (
	"context"
	"net/http"

	"github.com/elastic/go-elasticsearch/v8"

	"otus-recipe/app/api/parameters"
	"otus-recipe/app/config"
	"otus-recipe/app/models"
	db "otus-recipe/app/storage/db/sqlc"
	"otus-recipe/app/storage/elastic/elastic_index"
)

type Elastic interface {
	CreateIndex(ctx context.Context, index elastic_index.Index) error
	UpdateAliases(ctx context.Context, index elastic_index.Index) error
	BulkLoadRecipes(ctx context.Context, db db.Store, recipeIndex elastic_index.Index) error
	RefreshIndex(ctx context.Context) error
	SearchRecipes(ctx context.Context, params *parameters.RecipeListParams, paginated models.Paginated) (recipes []*elastic_index.Recipe, err error)
	GetRecipe(ctx context.Context, recipeId int64) (recipe *elastic_index.Recipe, err error)
}

type elastic struct {
	client *elasticsearch.Client
}

func New(config *config.Config) (Elastic, error) {
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.Elastic.URLS,
		Transport: &http.Transport{
			ResponseHeaderTimeout: config.Elastic.Timeout,
		},
	})
	if err != nil {
		return nil, err
	}

	return &elastic{client: es}, nil
}
