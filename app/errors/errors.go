package errors

import "errors"

var (
	RecipeNotFoundError   = errors.New("recipe not found")
	RecipeNotDeletedError = errors.New("recipe not deleted")
	RecipeVersionError    = errors.New("recipe version is not equal to yours")
)
