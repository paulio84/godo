package db

import (
	"github.com/paulio84/godo/internal/models/filter"
	"github.com/paulio84/godo/internal/models/todo"
)

type DBService interface {
	AddTodo(title string) error
	Connect() error
	Close() error
	EditTodo(id int, title string) error
	ListTodos(todoFilter filter.TodoFilter) ([]todo.Todo, error)
	PurgeTodos() error
	ToggleCompleted(id int) error
}
