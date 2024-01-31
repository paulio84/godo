package commands

import (
	"github.com/paulio84/godo/internal/models/core"
	"github.com/paulio84/godo/internal/services/database"
)

type PurgeCommand struct {
	dbService database.DBServicer
	result    core.DBResult
	parser    func(core.DBResult)
}

func NewPurgeCommand(db database.DBServicer, parser func(core.DBResult)) *PurgeCommand {
	return &PurgeCommand{
		dbService: db,
		parser:    parser,
	}
}

func (pc *PurgeCommand) Execute() {
	rowsAffected, err := pc.dbService.PurgeTodos()

	pc.result = core.DBResult{
		RowsAffected: rowsAffected,
		Err:          err,
	}

	pc.parser(pc.result)
}
