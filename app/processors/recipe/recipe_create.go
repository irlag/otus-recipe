package recipe

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	"otus-recipe/app/api/parameters"
	"otus-recipe/app/models"
	db "otus-recipe/app/storage/db/sqlc"
)

func (r *recipeProcessor) Create(ctx context.Context, params *parameters.RecipeCreateParams) (db.Recipe, error) {
	version, err := uuid.NewUUID()
	if err != nil {
		return db.Recipe{}, err
	}

	recipeFromRequest := db.CreateRecipeParams{
		Name: params.Name,
		Description: sql.NullString{
			String: params.Description,
			Valid:  true,
		},
		CookingTime: int16(params.CookingTime),
		Calories: sql.NullInt32{
			Int32: int32(params.Calories),
			Valid: true,
		},
		Proteins: sql.NullInt32{
			Int32: int32(params.Proteins),
			Valid: true,
		},
		Fats: sql.NullInt32{
			Int32: int32(params.Fats),
			Valid: true,
		},
		Carbohydrates: sql.NullInt32{
			Int32: int32(params.Carbohydrates),
			Valid: true,
		},
		Version: version,
	}

	recipe, err := r.store.CreateRecipe(ctx, recipeFromRequest)
	if err != nil {
		return db.Recipe{}, err
	}

	err = r.services.Notification.EventSend(ctx, models.RecipeUpdated{
		Name: recipe.Name,
	})
	if err != nil {
		return db.Recipe{}, err
	}

	return recipe, nil
}
