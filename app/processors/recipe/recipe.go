package recipe

import (
	"context"

	"otus-recipe/app/models"
	"otus-recipe/app/services"
	db "otus-recipe/app/storage/db/sqlc"
	"otus-recipe/app/storage/elastic"
	"otus-recipe/app/storage/elastic/elastic_index"

	"otus-recipe/app/api/parameters"
)

type Recipe interface {
	Get(ctx context.Context, id int64) (recipe *elastic_index.Recipe, err error)
	GetFromDb(ctx context.Context, id int64) (recipe db.Recipe, err error)
	List(ctx context.Context, paginated models.Paginated, params *parameters.RecipeListParams) (recipes []*elastic_index.Recipe, err error)
	ListCount(ctx context.Context) (int64, error)
	Create(ctx context.Context, params *parameters.RecipeCreateParams) (db.Recipe, error)
	Update(ctx context.Context, params *parameters.RecipeUpdateParams) (db.Recipe, error)
	Delete(ctx context.Context, params parameters.RecipeDeleteParams) error
}

type recipeProcessor struct {
	store         db.Store
	elasticsearch elastic.Elastic
	services      *services.Services
}

func NewRecipeProcessor(store db.Store, elasticsearch elastic.Elastic, services *services.Services) Recipe {
	return &recipeProcessor{
		store:         store,
		elasticsearch: elasticsearch,
		services:      services,
	}
}
