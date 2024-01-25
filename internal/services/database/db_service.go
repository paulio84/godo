package database

import (
	"github.com/paulio84/godo/internal/models/filter"
	"github.com/paulio84/godo/internal/models/todo"
)

type DBServicer interface {
	AddTodo(title string) (int, error)
	Init() error
	EditTodo(id int, title string) (int, error)
	ListTodos(todoFilter filter.TodoFilter) ([]todo.Todo, error)
	PurgeTodos() (int, error)
	ToggleCompleted(id int) error
}
