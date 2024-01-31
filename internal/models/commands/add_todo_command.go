package commands

import (
	"github.com/paulio84/godo/internal/models/core"
	"github.com/paulio84/godo/internal/services/database"
)

type AddTodoCommand struct {
	dbService database.DBServicer
	title     string
	result    core.DBResult
	parser    func(core.DBResult)
}

func NewAddTodoCommand(db database.DBServicer, parser func(core.DBResult), title string) *AddTodoCommand {
	return &AddTodoCommand{
		dbService: db,
		title:     title,
		parser:    parser,
	}
}

func (atc *AddTodoCommand) Execute() {
	rowsAffected, err := atc.dbService.AddTodo(atc.title)

	atc.result = core.DBResult{
		RowsAffected: rowsAffected,
		Err:          err,
	}

	atc.parser(atc.result)
}
