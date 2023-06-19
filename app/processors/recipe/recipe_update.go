package recipe

import (
	"context"
	"database/sql"

	"github.com/google/uuid"

	appErrors "otus-recipe/app/errors"

	"otus-recipe/app/api/parameters"
	db "otus-recipe/app/storage/db/sqlc"
)

func (r *recipeProcessor) Update(ctx context.Context, params *parameters.RecipeUpdateParams) (db.Recipe, error) {
	version, err := uuid.NewUUID()
	if err != nil {
		return db.Recipe{}, err
	}

	recipeFromRequest := db.UpdateRecipeParams{
		ID:   params.RecipeID,
		Name: params.Name,
		Description: sql.NullString{
			String: params.Description,
			Valid:  true,
		},
		CookingTime: 0,
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
	tx, err := r.store.Begin()
	if err != nil {
		return db.Recipe{}, err
	}
	defer tx.Rollback()

	qtx := r.store.WithTx(tx)

	currentRecipe, err := qtx.GetRecipeForUpdate(ctx, recipeFromRequest.ID)
	if err != nil {
		return db.Recipe{}, err
	}

	if currentRecipe.Version.String() != params.Version {
		return db.Recipe{}, appErrors.RecipeVersionError
	}

	recipe, err := qtx.UpdateRecipe(ctx, recipeFromRequest)
	if err != nil {
		return db.Recipe{}, err
	}

	err = tx.Commit()
	if err != nil {
		return db.Recipe{}, err
	}

	return recipe, nil
}
