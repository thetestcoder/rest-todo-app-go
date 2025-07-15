package storage

import (
	"github.com/thetestcoder/rest-todo/internal/models"
	"github.com/thetestcoder/rest-todo/internal/storage"
	"os"
	"testing"
)

func TestFileStorage(t *testing.T) {
	tempFile := "test_todos.json"
	fs := storage.NewFileStorage(tempFile)
	defer os.Remove(tempFile)

	err := fs.Initiate()

	if err != nil {
		t.Errorf("error to initiate file storage: %v", err)
	}

	t.Run("Test Store", func(t *testing.T) {
		todo := models.TODO{
			ID:          1,
			Title:       "Test",
			Description: "Test",
			Completed:   false,
		}

		_ = fs.Store(todo)

		todos, _ := fs.List()
		if len(todos) != 1 {
			t.Errorf("Expeceted 1 todo but got %d", len(todos))
		}

	})

	t.Run("Test Get Single", func(t *testing.T) {
		todo, _ := fs.GetSingle(1)
		if todo.ID != 1 {
			t.Errorf("Expeceted TODO ID: 1 but got %d", todo.ID)
		}
	})

	t.Run("Test List", func(t *testing.T) {
		todos, _ := fs.List()
		if len(todos) != 1 {
			t.Errorf("Expeceted 1 todo but got %d", len(todos))
		}
	})

	t.Run("Test Update TODO", func(t *testing.T) {
		todo, _ := fs.GetSingle(1)

		todo.Title = "Updated"
		todo.Completed = true
		fs.Update(todo)

		todo, _ = fs.GetSingle(1)

		if todo.Completed != true {
			t.Errorf("Update is not happening")
		}
	})

	t.Run("Test Delete TODO", func(t *testing.T) {
		todo, _ := fs.GetSingle(1)
		fs.Delete(todo)

		todos, _ := fs.List()

		if len(todos) != 0 {
			t.Errorf("Excepted 0 todos but got %d", len(todos))
		}
	})

}
