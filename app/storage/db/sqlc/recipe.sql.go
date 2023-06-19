// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: recipe.sql

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const CreateRecipe = `-- name: CreateRecipe :one
INSERT INTO recipe (
    name,
    description,
    cooking_time,
    calories,
    proteins,
    fats,
    carbohydrates,
    version
) VALUES (
     $1, $2, $3, $4, $5, $6, $7, $8
 ) RETURNING id, name, description, cooking_time, calories, proteins, fats, carbohydrates, rating, vote_count, vote_sum, version
`

type CreateRecipeParams struct {
	Name          string         `db:"name" json:"name"`
	Description   sql.NullString `db:"description" json:"description"`
	CookingTime   int16          `db:"cooking_time" json:"cooking_time"`
	Calories      sql.NullInt32  `db:"calories" json:"calories"`
	Proteins      sql.NullInt32  `db:"proteins" json:"proteins"`
	Fats          sql.NullInt32  `db:"fats" json:"fats"`
	Carbohydrates sql.NullInt32  `db:"carbohydrates" json:"carbohydrates"`
	Version       uuid.UUID      `db:"version" json:"version"`
}

func (q *Queries) CreateRecipe(ctx context.Context, arg CreateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, CreateRecipe,
		arg.Name,
		arg.Description,
		arg.CookingTime,
		arg.Calories,
		arg.Proteins,
		arg.Fats,
		arg.Carbohydrates,
		arg.Version,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CookingTime,
		&i.Calories,
		&i.Proteins,
		&i.Fats,
		&i.Carbohydrates,
		&i.Rating,
		&i.VoteCount,
		&i.VoteSum,
		&i.Version,
	)
	return i, err
}

const DeleteRecipe = `-- name: DeleteRecipe :execrows
DELETE FROM recipe
WHERE id = $1
`

func (q *Queries) DeleteRecipe(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, DeleteRecipe, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const GetRecipe = `-- name: GetRecipe :one
SELECT id, name, description, cooking_time, calories, proteins, fats, carbohydrates, rating, vote_count, vote_sum, version FROM recipe
WHERE id = $1
`

func (q *Queries) GetRecipe(ctx context.Context, id int64) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, GetRecipe, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CookingTime,
		&i.Calories,
		&i.Proteins,
		&i.Fats,
		&i.Carbohydrates,
		&i.Rating,
		&i.VoteCount,
		&i.VoteSum,
		&i.Version,
	)
	return i, err
}

const GetRecipeForUpdate = `-- name: GetRecipeForUpdate :one
SELECT id, name, description, cooking_time, calories, proteins, fats, carbohydrates, rating, vote_count, vote_sum, version FROM recipe
WHERE id = $1 LIMIT 1
FOR UPDATE
`

func (q *Queries) GetRecipeForUpdate(ctx context.Context, id int64) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, GetRecipeForUpdate, id)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CookingTime,
		&i.Calories,
		&i.Proteins,
		&i.Fats,
		&i.Carbohydrates,
		&i.Rating,
		&i.VoteCount,
		&i.VoteSum,
		&i.Version,
	)
	return i, err
}

const ListRecipe = `-- name: ListRecipe :many
SELECT id, name, description, cooking_time, calories, proteins, fats, carbohydrates, rating, vote_count, vote_sum, version FROM recipe
ORDER BY id DESC
LIMIT $1
OFFSET $2
`

type ListRecipeParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) ListRecipe(ctx context.Context, arg ListRecipeParams) ([]Recipe, error) {
	rows, err := q.db.QueryContext(ctx, ListRecipe, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Recipe{}
	for rows.Next() {
		var i Recipe
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CookingTime,
			&i.Calories,
			&i.Proteins,
			&i.Fats,
			&i.Carbohydrates,
			&i.Rating,
			&i.VoteCount,
			&i.VoteSum,
			&i.Version,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const ListRecipeCount = `-- name: ListRecipeCount :one
SELECT count(*) FROM recipe
`

func (q *Queries) ListRecipeCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, ListRecipeCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const UpdateRecipe = `-- name: UpdateRecipe :one
UPDATE recipe SET
    name = $2,
    description = $3,
    cooking_time = $4,
    calories = $5,
    proteins = $6,
    fats = $7,
    carbohydrates = $8,
    version = $9
WHERE id = $1 RETURNING id, name, description, cooking_time, calories, proteins, fats, carbohydrates, rating, vote_count, vote_sum, version
`

type UpdateRecipeParams struct {
	ID            int64          `db:"id" json:"id"`
	Name          string         `db:"name" json:"name"`
	Description   sql.NullString `db:"description" json:"description"`
	CookingTime   int16          `db:"cooking_time" json:"cooking_time"`
	Calories      sql.NullInt32  `db:"calories" json:"calories"`
	Proteins      sql.NullInt32  `db:"proteins" json:"proteins"`
	Fats          sql.NullInt32  `db:"fats" json:"fats"`
	Carbohydrates sql.NullInt32  `db:"carbohydrates" json:"carbohydrates"`
	Version       uuid.UUID      `db:"version" json:"version"`
}

func (q *Queries) UpdateRecipe(ctx context.Context, arg UpdateRecipeParams) (Recipe, error) {
	row := q.db.QueryRowContext(ctx, UpdateRecipe,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CookingTime,
		arg.Calories,
		arg.Proteins,
		arg.Fats,
		arg.Carbohydrates,
		arg.Version,
	)
	var i Recipe
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CookingTime,
		&i.Calories,
		&i.Proteins,
		&i.Fats,
		&i.Carbohydrates,
		&i.Rating,
		&i.VoteCount,
		&i.VoteSum,
		&i.Version,
	)
	return i, err
}
