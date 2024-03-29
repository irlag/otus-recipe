// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"context"
)

type Querier interface {
	CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error)
	DeleteRecipe(ctx context.Context, id int64) (int64, error)
	GetRecipe(ctx context.Context, id int64) (Recipe, error)
	GetRecipeForUpdate(ctx context.Context, id int64) (Recipe, error)
	ListAllRecipes(ctx context.Context) ([]Recipe, error)
	ListRecipe(ctx context.Context, arg ListRecipeParams) ([]Recipe, error)
	ListRecipeCount(ctx context.Context) (int64, error)
	UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) (Recipe, error)
}

var _ Querier = (*Queries)(nil)
