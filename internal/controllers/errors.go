package controllers

import "errors"

var (
	ErrNoFound             = errors.New("Not found")
	ErrInternalServerError = errors.New("Internal server error")
	ErrDuplicateKey        = errors.New("Duplicate key")
)
