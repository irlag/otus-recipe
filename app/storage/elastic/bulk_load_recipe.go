package elastic

import (
	"context"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esutil"

	"otus-recipe/app/storage/elastic/elastic_index"

	db "otus-recipe/app/storage/db/sqlc"
)

func (e *elastic) BulkLoadRecipes(ctx context.Context, db db.Store, recipeIndex elastic_index.Index) error {
	indexer, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
		Index:      recipeIndex.GetName(),
		Client:     e.client,
		NumWorkers: 2,
		FlushBytes: 1e+7,
	})
	if err != nil {
		return err
	}

	recipes, err := db.ListAllRecipes(ctx)
	if err != nil {
		return err
	}
	for _, recipe := range recipes {
		model := elastic_index.Recipe{}

		esRecipe := model.GetFromDBModelMapping(recipe)
		body, err := esRecipe.MarshalJSON()
		if err != nil {
			return err
		}
		err = indexer.Add(
			ctx,
			esutil.BulkIndexerItem{
				Action:     "index",
				DocumentID: model.GetDocumentId(),
				Body:       strings.NewReader(string(body)),
			})
		if err != nil {
			return err
		}
	}

	err = indexer.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}
