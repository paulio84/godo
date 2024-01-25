package main

import (
	"database/sql"

	"github.com/paulio84/godo/internal/models/commands"
	"github.com/paulio84/godo/internal/services/database"
)

type cli struct {
	dbService database.DBServicer
	command   commands.Commander
}

func newCLI(db *sql.DB) cli {
	return cli{
		dbService: database.NewSQLiteService(db),
	}
}
