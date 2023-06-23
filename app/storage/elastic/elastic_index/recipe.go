package elastic_index

//go:generate easyjson

import (
	"strconv"
	"time"

	"otus-recipe/app/builders"
	db "otus-recipe/app/storage/db/sqlc"
)

const (
	RecipeIndexAlias = "recipe"
)

//easyjson:json
type Recipe struct {
	ID            string `db:"id" json:"id"`
	Name          string `db:"name" json:"name"`
	Description   string `db:"description" json:"description"`
	CookingTime   int16  `db:"cooking_time" json:"cooking_time"`
	Calories      *int32 `db:"calories" json:"calories"`
	Proteins      *int32 `db:"proteins" json:"proteins"`
	Fats          *int32 `db:"fats" json:"fats"`
	Carbohydrates *int32 `db:"carbohydrates" json:"carbohydrates"`
	Version       string `db:"version" json:"version"`
}

func (r *Recipe) GetDocumentId() string {
	return r.ID
}

func (r *Recipe) GetTypeMapping() string {
	return RecipeIndexAlias
}

func (r *Recipe) GetFromDBModelMapping(recipe db.Recipe) Recipe {
	return Recipe{
		ID:            strconv.Itoa(int(recipe.ID)),
		Name:          recipe.Name,
		Description:   recipe.Description.String,
		CookingTime:   recipe.CookingTime,
		Calories:      builders.GetIntValueFromSqlNull(recipe.Calories),
		Proteins:      builders.GetIntValueFromSqlNull(recipe.Proteins),
		Fats:          builders.GetIntValueFromSqlNull(recipe.Fats),
		Carbohydrates: builders.GetIntValueFromSqlNull(recipe.Carbohydrates),
		Version:       recipe.Version.String(),
	}
}

func NewRecipe(indexName string) Index {
	return &esIndex{
		Name:  indexName + "-" + time.Now().Format(indexNameTime),
		Alias: indexName,
	}
}
