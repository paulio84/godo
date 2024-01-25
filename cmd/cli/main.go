package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/paulio84/godo/internal/models/commands"
	"github.com/paulio84/godo/internal/models/filter"
)

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cli := newCLI(db)
	cli.dbService.Init()

	// get and process command-line args passed to the program
	if len(os.Args) == 1 {
		os.Args = append(os.Args, "-h")
	}

	// DEBUGGING
	// a, b := cli.dbService.AddTodo("A todo item")
	// fmt.Println(a, b)
	// cli.dbService.ToggleCompleted(8)
	// l, _ := cli.dbService.ListTodos(filter.All)
	// fmt.Println(l)
	// cli.dbService.EditTodo(1, "A todo itemaaa")
	// cli.dbService.PurgeTodos()

	switch os.Args[1] {
	case "-l", "list":
		cli.command = commands.NewListCommand(cli.dbService, displayTodoData, filter.NotCompleted)
	case "-la", "list-all":
		cli.command = commands.NewListCommand(cli.dbService, displayTodoData, filter.All)
	case "-ld", "list-done":
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
