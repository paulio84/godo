package commands

import (
	"github.com/paulio84/godo/internal/models/core"
	"github.com/paulio84/godo/internal/services/database"
)

type ToggleDoneCommand struct {
	dbService database.DBServicer
	result    core.DBResult
	parser    func(core.DBResult)
	id        int
}

func NewToggleDoneCommand(db database.DBServicer, parser func(core.DBResult), id int) *ToggleDoneCommand {
	return &ToggleDoneCommand{
		dbService: db,
		parser:    parser,
		id:        id,
	}
}

func (tdc *ToggleDoneCommand) Execute() {
	rowsAffected, err := tdc.dbService.ToggleCompleted(tdc.id)

	tdc.result = core.DBResult{
		RowsAffected: rowsAffected,
		Err:          err,
	}

	tdc.parser(tdc.result)
}
