package db

import "errors"

var (
	ErrNoUserFound error = errors.New("no user found")
)