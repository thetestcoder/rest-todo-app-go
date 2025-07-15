package storage

import (
	"encoding/json"
	"fmt"
	"github.com/thetestcoder/rest-todo/internal/models"
	"os"
	"sync"
)

type FileStorage struct {
	filename string
	mu       sync.Mutex
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{
		filename: filename,
	}
}

func (fs *FileStorage) List() ([]models.TODO, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	todos, err := fs.readTodos()

	if err != nil {
		return nil, fmt.Errorf("error to reading task: %w", err)
	}

	return todos, nil
}

func (fs *FileStorage) Store(todo models.TODO) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	todos, _ := fs.readTodos()
	todo.ID = fs.generateID(todos)
	todos = append(todos, todo)
	err := fs.writeTodos(todos)

	if err != nil {
		return fmt.Errorf("error to store task: %w", err)
	}

	return nil
}

func (fs *FileStorage) GetSingle(id int64) (*models.TODO, error) {

	fs.mu.Lock()
	defer fs.mu.Unlock()

	todos, _ := fs.readTodos()
	for _, todo := range todos {
		if todo.ID == id {
			return &todo, nil
		}
	}

	return nil, fmt.Errorf("error to get single task ID: %d", id)
}

func (fs *FileStorage) Update(todo *models.TODO) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	todos, _ := fs.readTodos()
	for index, singleTodo := range todos {
		if singleTodo.ID == todo.ID {
			todos[index] = *todo
		}
	}

	err := fs.writeTodos(todos)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileStorage) Delete(todo *models.TODO) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	todos, _ := fs.readTodos()
	for index, singleTodo := range todos {
		if todo.ID == singleTodo.ID {
			todos = append(todos[:index], todos[index+1:]...)
		}
	}

	err := fs.writeTodos(todos)
	if err != nil {
		return err
	}

	return nil

}

func (fs *FileStorage) Initiate() error {
	if _, err := os.Stat(fs.filename); os.IsNotExist(err) {
		return os.WriteFile(fs.filename, []byte("[]"), 0644)
	}
	return nil
}

func (fs *FileStorage) readTodos() ([]models.TODO, error) {
	data, err := os.ReadFile(fs.filename)

	if err != nil {
		return nil, fmt.Errorf("error to read file: %w", err)
	}

	var todos []models.TODO
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, fmt.Errorf("error to unmarshal file: %w", err)
	}

	return todos, nil
}

func (fs *FileStorage) writeTodos(todos []models.TODO) error {
	data, err := json.Marshal(todos)
	if err != nil {
		return fmt.Errorf("error to marshal file: %w", err)
	}

	return os.WriteFile(fs.filename, data, 0644)
}

func (fs *FileStorage) generateID(todos []models.TODO) int64 {
	if len(todos) == 0 {
		return 1
	}

	return todos[len(todos)-1].ID + 1
}
