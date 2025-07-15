package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thetestcoder/rest-todo/internal/models"
	"github.com/thetestcoder/rest-todo/internal/storage"
	"net/http"
	"strconv"
)

type TodoHandler struct {
	storage *storage.FileStorage
}

func NewTodoHandler(fileStorage *storage.FileStorage) *TodoHandler {
	return &TodoHandler{
		storage: fileStorage,
	}
}

func (handler *TodoHandler) ListTodo(writer http.ResponseWriter, request *http.Request) {
	todos, err := handler.storage.List()

	if err != nil {
		panic(err)
	}

	json.NewEncoder(writer).Encode(todos)
}

func (handler *TodoHandler) StoreTodo(writer http.ResponseWriter, request *http.Request) {
	var newTodo models.TODO

	if err := json.NewDecoder(request.Body).Decode(&newTodo); err != nil {
		panic(err)
	}

	handler.storage.Store(newTodo)

	fmt.Fprint(writer, "Successfully stored")
}

func (handler *TodoHandler) GetSingleTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	todo, err := handler.storage.GetSingle(id)
	if err != nil {
		fmt.Errorf("something went wrong. %w", err)
	}

	json.NewEncoder(writer).Encode(todo)
}

func (handler *TodoHandler) UpdateTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	todo, err := handler.storage.GetSingle(id)

	if err != nil {
		fmt.Errorf("something went wrong: %w", err)
	}

	if err := json.NewDecoder(request.Body).Decode(&todo); err != nil {
		panic(err)
	}
	err = handler.storage.Update(todo)
	if err != nil {
		fmt.Errorf("something went wrong: %w", err)
	}

	fmt.Fprint(writer, "Successfully updated")
}

func (handler *TodoHandler) DeleteTodo(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.ParseInt(params["id"], 10, 64)

	todo, _ := handler.storage.GetSingle(id)

	err := handler.storage.Delete(todo)
	if err != nil {
		fmt.Errorf("something went wrong: %w", err)
	}

	fmt.Fprint(writer, "Successfully deleted")

}
