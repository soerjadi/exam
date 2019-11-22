package models

import "errors"

var (
	// ErrNotFound will throw if models not exists
	ErrNotFound = errors.New("Not found")

	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
)
