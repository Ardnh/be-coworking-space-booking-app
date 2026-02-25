package errors

import "errors"

var (
	ErrNotFound       = errors.New("resource not found")
	ErrConflict       = errors.New("resource already exists")
	ErrInternalServer = errors.New("internal server error")
	ErrInvalidInput   = errors.New("invalid input")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
)

var (
	ErrImageTooLarge      = errors.New("image too large")
	ErrImageInvalidFormat = errors.New("image format not supported")
	ErrImageInvalidSize   = errors.New("image size not supported")
)
