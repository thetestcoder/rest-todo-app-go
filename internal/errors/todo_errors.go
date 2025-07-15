package errors

import "errors"

var (
	ErrEmptyTitle         = errors.New("title is required")
	ErrTitleTooLong       = errors.New("title cannot be more than 255 characters")
	ErrDescriptionTooLong = errors.New("description cannot be more than 500 characters")
)
