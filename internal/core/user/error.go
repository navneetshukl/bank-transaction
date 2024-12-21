package user

import "errors"

var (
	ErrInvalidUser             error = errors.New("invalid user")
	ErrGeneratingAccountNumber error = errors.New("failed to generate account number")
)
