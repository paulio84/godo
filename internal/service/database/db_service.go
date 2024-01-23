package database

import (
	"github.com/paulio84/godo/internal/models/filter"
	"github.com/paulio84/godo/internal/models/todo"
)

type DBServicer interface {
	AddTodo(title string) error
	Init() error
	EditTodo(id int, title string) error
	ListTodos(todoFilter filter.TodoFilter) ([]todo.Todo, error)
	PurgeTodos() error
	ToggleCompleted(id int) error
}
