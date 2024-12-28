package user

import "errors"

var (
	ErrInvalidUser             error = errors.New("invalid user")
	ErrGeneratingAccountNumber error = errors.New("failed to generate account number")
	ErrCreatingUser            error = errors.New("failed to create user")
	ErrGettingAmount           error = errors.New("failed to get amount")
	ErrUpdatingAmount          error = errors.New("failed to update amount")
)
