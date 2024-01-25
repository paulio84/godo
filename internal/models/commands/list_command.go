package commands

import (
	"github.com/paulio84/godo/internal/models/core"
	"github.com/paulio84/godo/internal/models/filter"
	"github.com/paulio84/godo/internal/models/todo"
	"github.com/paulio84/godo/internal/services/database"
)

type ListCommand struct {
	dbService  database.DBServicer
	listFilter filter.TodoFilter
	result     core.DBResult
	parser     func(core.DBResult)
}

func NewListCommand(db database.DBServicer, parser func(core.DBResult), tf filter.TodoFilter) *ListCommand {
	return &ListCommand{
		dbService:  db,
		parser:     parser,
		listFilter: tf,
	}
}

func (lc *ListCommand) Execute() {
	data, err := lc.dbService.ListTodos(lc.listFilter)
	if err != nil {
		lc.result = core.DBResult{
			RowsAffected: 0,
			Data:         []todo.Todo{},
			Err:          err,
		}
	}

	lc.result = core.DBResult{
		RowsAffected: 0,
		Data:         data,
	}
}

func (lc ListCommand) ParseOutput() {
	lc.parser(lc.result)
}
