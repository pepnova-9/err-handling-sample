package usecase

import "errors"

var (
	ErrUserNotFound = errors.New("user not found")
	ErrUnauthorized = errors.New("unauthorized")

	ErrUnexpectedError = errors.New("unexpected error")
)
