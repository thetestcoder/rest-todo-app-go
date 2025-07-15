package models

import "github.com/thetestcoder/rest-todo/internal/errors"

type TODO struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func (todo *TODO) Validate() error {
	if todo.Title == "" {
		return errors.ErrEmptyTitle
	}

	if len(todo.Title) > 100 {
		return errors.ErrTitleTooLong
	}

	if len(todo.Description) > 500 {
		return errors.ErrDescriptionTooLong
	}

	return nil

}
