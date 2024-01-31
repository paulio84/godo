package commands

import (
	"github.com/paulio84/godo/internal/models/core"
	"github.com/paulio84/godo/internal/services/database"
)

type EditTodoCommand struct {
	dbService database.DBServicer
	result    core.DBResult
	parser    func(core.DBResult)
	id        int
	title     string
}

func NewEditTodoCommand(db database.DBServicer, parser func(core.DBResult), id int, title string) *EditTodoCommand {
	return &EditTodoCommand{
		dbService: db,
		parser:    parser,
		id:        id,
		title:     title,
	}
}

func (etc *EditTodoCommand) Execute() {
	rowsAffected, err := etc.dbService.EditTodo(etc.id, etc.title)

	etc.result = core.DBResult{
		RowsAffected: rowsAffected,
		Err:          err,
	}

	etc.parser(etc.result)
}
