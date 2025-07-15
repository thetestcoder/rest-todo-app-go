# Developer Hints

## Project Overview
Implement a To-Do List RESTful API with CRUD operations using file I/O for persistence.

## Technical Requirements
- Use Gorilla Mux for routing (`github.com/gorilla/mux`)
- Implement JSON encoding/decoding for request/response handling
- Store data in a JSON file on disk
- Follow RESTful API principles
- Implement proper error handling

## Project Structure
```
├── cmd/
│   └── rest/          # Application entrypoint
│       └── main.go
├── internal/
│   ├── api/          # HTTP router and server logic
│   ├── handlers/     # HTTP handlers
│   ├── models/       # Data models
│   ├── errors/       # errors constants
│   ├── middleware/   # middleware logics
│   ├── storage/      # Where todo data is stored
└── tests/            # Integration tests
```

## Implementation Tips
1. Start with defining the Todo struct in models
2. Implement the repository layer with file I/O operations
3. Create handlers for each CRUD operation
4. Set up routes in the main package
5. Add middleware for logging and error handling
6. Implement proper validation for inputs

## Testing
Use Go's testing package to write unit tests for each component.
Implement integration tests to verify the complete API functionality.