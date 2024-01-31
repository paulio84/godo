package main

import (
	"bufio"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/paulio84/godo/internal/models/commands"
	"github.com/paulio84/godo/internal/models/core"
	"github.com/paulio84/godo/internal/models/filter"
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

func (c *cli) initialiseDB() {
	c.dbService.Init()
}

func (c *cli) addTodoHandler(s string) error {
	if s == "" {
		return errors.New(core.TitleMandatory)
	}

	c.command = commands.NewAddTodoCommand(c.dbService, displayCreated, s)
	return nil
}

func (c *cli) editTodoHandler(s string) error {
	// parse the todo id
	id, err := strconv.Atoi(s)
	if err != nil {
		return errors.New(core.CannotParseTodoID)
	}

	// get the new title from the user
	var newTitle string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a new title -> ")
	scanner.Scan()
	newTitle = scanner.Text()
	if newTitle == "" {
		return errors.New(core.TitleMandatory)
	}

	c.command = commands.NewEditTodoCommand(c.dbService, displayEdited, id, newTitle)
	return nil
}

func (c *cli) purgeTodosHandler(s string) error {
	c.command = commands.NewPurgeCommand(c.dbService, displayPurged)
	return nil
}

func (c *cli) toggleTodoHandler(s string) error {
	id, err := strconv.Atoi(s)
	if err != nil {
		return errors.New(core.CannotParseTodoID)
	}

	c.command = commands.NewToggleDoneCommand(c.dbService, displayToggled, id)
	return nil
}

func (c *cli) listTodosHandler(s string) error {
	c.command = commands.NewListCommand(c.dbService, displayTodoData, filter.NotCompleted)
	return nil
}

func (c *cli) listDoneTodosHandler(s string) error {
	c.command = commands.NewListCommand(c.dbService, displayTodoData, filter.OnlyCompleted)
	return nil
}

func (c *cli) listAllTodosHandler(s string) error {
	c.command = commands.NewListCommand(c.dbService, displayTodoData, filter.All)
	return nil
}
