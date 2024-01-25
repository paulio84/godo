package core

import "github.com/paulio84/godo/internal/models/todo"

type DBResult struct {
	RowsAffected int
	Data         []todo.Todo
	Err          error
}
