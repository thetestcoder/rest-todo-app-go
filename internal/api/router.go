package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thetestcoder/rest-todo/configs"
	"github.com/thetestcoder/rest-todo/internal/handler"
	"github.com/thetestcoder/rest-todo/internal/middleware"
	"github.com/thetestcoder/rest-todo/internal/storage"
	"net/http"
	"os"
)

// setupRouter initializes and configures the HTTP router for the application.
// It sets up the routes and corresponding handlers for various endpoints, including CRUD operations for todos.
// Returns an http.Handler that handles the incoming HTTP requests.
func setupRouter() http.Handler {
	router := mux.NewRouter()

	//
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "API Working fine")
	}).Methods("GET")

	config := configs.NewConfig("./configs/config.json")
	if config == nil {
		fmt.Println("Error to load config")
		os.Exit(1)
	}

	// setup file
	todoStorage := storage.NewFileStorage(config.Storage.FilePath)
	if err := todoStorage.Initiate(); err != nil {
		fmt.Printf("Failed to initialize storage: %v\n", err)
		os.Exit(1)
	}
	todoHandler := handler.NewTodoHandler(todoStorage)

	router.HandleFunc("/todos", todoHandler.ListTodo).Methods("GET")
	router.HandleFunc("/todos", todoHandler.StoreTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", todoHandler.GetSingleTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PATCH")
	router.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")

	router.Use(middleware.LoggingMiddleware, middleware.RestMiddleware)

	return router
}
