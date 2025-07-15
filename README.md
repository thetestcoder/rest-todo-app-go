# To-Do List RESTful API

A simple RESTful API for managing to-do items with CRUD operations using Go.

## Requirements

- Go 1.24+
- Gorilla Mux for routing
- JSON for data serialization/deserialization
- File I/O for persistence

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/todo-api.git
cd todo-api

# Install dependencies
make deps
```

## Usage

```bash
# Build the application
make build

# Run the application
make run

# Run tests
make test
```

## API Endpoints

- `GET /todos` - Get all to-do items
- `GET /todos/{id}` - Get a specific to-do item
- `POST /todos` - Create a new to-do item
- `PUT /todos/{id}` - Update a to-do item
- `DELETE /todos/{id}` - Delete a to-do item

## Data Structure

```json
{
  "id": int64,
  "title": "string",
  "description": "string",
  "completed": boolean,
}
```

## Development

```bash
# Format code
make fmt

# Run linter
make lint

# Clean up build artifacts
make clean
```

## License

MIT