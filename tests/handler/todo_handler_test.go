package handler

import (
	"bytes"
	"encoding/json"
	"github.com/thetestcoder/rest-todo/internal/handler"
	"github.com/thetestcoder/rest-todo/internal/models"
	"github.com/thetestcoder/rest-todo/internal/storage"
	"net/http/httptest"
	"os"
	"testing"
)

func setupTestHandler() (*handler.TodoHandler, string) {
	tempFile := "test_todos.json"
	fs := storage.NewFileStorage(tempFile)
	fs.Initiate()
	return handler.NewTodoHandler(fs), tempFile
}

func TestTodoHandler_Integration(t *testing.T) {
	handler, tempFile := setupTestHandler()
	defer os.Remove(tempFile)

	t.Run("Store TODO", func(t *testing.T) {
		todo := models.TODO{
			Title:       "Test",
			Description: "Test",
			Completed:   false,
		}

		todoJSON, _ := json.Marshal(todo)

		req := httptest.NewRequest("POST", "/todos", bytes.NewBuffer(todoJSON))
		w := httptest.NewRecorder()

		handler.StoreTodo(w, req)

		if w.Code != 200 {
			t.Errorf("Expected 200 but got %d", w.Code)
		}
		if w.Body.String() != "Successfully stored" {
			t.Errorf("Expected Successfully stored but got %s", w.Body.String())
		}

	})

	t.Run("List TODO", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/todos", nil)
		w := httptest.NewRecorder()
		handler.ListTodo(w, req)

		var todos []models.TODO
		json.NewDecoder(w.Body).Decode(&todos)

		if len(todos) != 1 {
			t.Errorf("Excepted 1 todo but got %d", len(todos))
		}
	})

	t.Run("Single TODO", func(t *testing.T) {
		panic("implement me")
	})

	t.Run("Update TODO", func(t *testing.T) {
		panic("implement me")
	})

	t.Run("Delete TODO", func(t *testing.T) {
		panic("implement me")
	})

}
