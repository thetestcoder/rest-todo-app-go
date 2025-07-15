package storage

import "github.com/thetestcoder/rest-todo/internal/models"

type Storage interface {
	List() ([]models.TODO, error)
	Store(todo *models.TODO) error
	GetSingle(id int64) (models.TODO, error)
	Update(todo models.TODO) error
	Delete(todo models.TODO) error
}
