package models

import "errors"

var (
	// ErrNotFound will throw if models not exists
	ErrNotFound = errors.New("Not found")
)
