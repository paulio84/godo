package main

import (
	"bufio"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/paulio84/godo/internal/models/commands"
	"github.com/paulio84/godo/internal/models/core"
	"github.com/paulio84/godo/internal/models/filter"
)

func main() {
	defer errorHandler()

	// check if the database file exists, if it does not then we need to create
	// it and create the table `todo`
	var shouldCreateDB bool
	f, err := os.Open("todo.db")
	if err != nil {
		shouldCreateDB = true
	}
	f.Close()

	// get a connection pool to the database
	db, err := openDB()
	if err != nil {
		panic("error: " + core.DBConnectionError)
	}
	defer db.Close()

	cli := newCLI(db)
	if shouldCreateDB {
		cli.dbService.Init()
	}

	// setup flags to be used by the command-line
	flag.Func("add", "`<todo item>` to be added.", func(s string) error {
		if s == "" {
			return errors.New(core.TitleMandatory)
		}

		cli.command = commands.NewAddTodoCommand(cli.dbService, displayCreated, s)
		return nil
	})

	flag.Func("edit", "`<todo id>` to be updated.", func(s string) error {
		// parse the todo id
		id, err := strconv.Atoi(s)
		if err != nil {
			return errors.New(core.CannotParseTodoId)
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

		cli.command = commands.NewEditTodoCommand(cli.dbService, displayEdited, id, newTitle)
		return nil
	})

	flag.BoolFunc("purge", "Remove completed todo's.", func(s string) error {
		cli.command = commands.NewPurgeCommand(cli.dbService, displayPurged)
		return nil
	})

	flag.Func("toggle", "`<todo id>` Todo to be completed, or uncompleted.", func(s string) error {
		id, err := strconv.Atoi(s)
		if err != nil {
			return errors.New(core.CannotParseTodoId)
		}

		cli.command = commands.NewToggleDoneCommand(cli.dbService, displayToggled, id)
		return nil
	})

	flag.BoolFunc("list", "list incomplete todo's.", func(s string) error {
		cli.command = commands.NewListCommand(cli.dbService, displayTodoData, filter.NotCompleted)
		return nil
	})

	flag.BoolFunc("list-done", "list completed todo's.", func(s string) error {
		cli.command = commands.NewListCommand(cli.dbService, displayTodoData, filter.OnlyCompleted)

		return nil
	})

	flag.BoolFunc("list-all", "list all todo's.", func(s string) error {
		cli.command = commands.NewListCommand(cli.dbService, displayTodoData, filter.All)

		return nil
	})

	flag.Parse()

	cli.command.Execute()
}

func openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "todo.db")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func errorHandler() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
