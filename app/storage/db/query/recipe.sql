-- name: GetRecipe :one
SELECT * FROM recipe
WHERE id = $1;

-- name: GetRecipeForUpdate :one
SELECT * FROM recipe
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: ListRecipe :many
SELECT * FROM recipe
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: ListRecipeCount :one
SELECT count(*) FROM recipe;

-- name: CreateRecipe :one
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
 ) RETURNING *;

-- name: UpdateRecipe :one
UPDATE recipe SET
    name = $2,
    description = $3,
    cooking_time = $4,
    calories = $5,
    proteins = $6,
    fats = $7,
    carbohydrates = $8,
    version = $9
WHERE id = $1 RETURNING *;

-- name: DeleteRecipe :execrows
DELETE FROM recipe
WHERE id = $1;