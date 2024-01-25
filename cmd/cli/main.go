package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/paulio84/godo/internal/models/commands"
	"github.com/paulio84/godo/internal/models/core"
	"github.com/paulio84/godo/internal/models/filter"
)

func main() {
	defer errorHandler()

	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cli := newCLI(db)
	cli.dbService.Init()

	// get and process command-line args passed to the program
	args := processArgs(os.Args)

	switch args.tag {
	case core.A, core.Add:
		cli.command = commands.NewAddTodoCommand(cli.dbService, displayCreated, args.todoTitle)
	case core.L, core.List:
		cli.command = commands.NewListCommand(cli.dbService, displayTodoData, filter.NotCompleted)
	case core.LA, core.ListAll:
		cli.command = commands.NewListCommand(cli.dbService, displayTodoData, filter.All)
	case core.LD, core.ListDone:
		cli.command = commands.NewListCommand(cli.dbService, displayTodoData, filter.OnlyCompleted)
	default: // unknown, -h or help
		cli.command = commands.NewHelpCommand(displayHelpText)
	}

	cli.command.Execute()
	cli.command.ParseOutput()
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

type cliArgs struct {
	tag       string
	id        int
	todoTitle string
}

func processArgs(osArgs []string) (args cliArgs) {
	// fmt.Println("ARGS->", osArgs)
	if len(osArgs) == 1 {
		args = cliArgs{
			tag: core.H,
		}
		return
	}

	args = cliArgs{
		tag: osArgs[1],
	}

	switch args.tag {
	case core.L, core.List:
	case core.LA, core.ListAll:
	case core.LD, core.ListDone:
		return
	case core.A, core.Add:
		if len(osArgs) <= 2 {
			panic(core.TitleMandatory)
		}
		args.todoTitle = osArgs[2]
	}

	return
}

func errorHandler() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
