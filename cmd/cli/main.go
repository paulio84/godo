package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/paulio84/godo/internal/models/core"
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
		cli.initialiseDB()
	}

	flag.Func("add", "`<todo item>` to be added.", cli.addTodoHandler)
	flag.Func("edit", "`<todo id>` to be updated.", cli.editTodoHandler)
	flag.BoolFunc("purge", "Remove completed todo's.", cli.purgeTodosHandler)
	flag.Func("toggle", "`<todo id>` Todo to be completed, or uncompleted.", cli.toggleTodoHandler)
	flag.BoolFunc("list", "list incomplete todo's.", cli.listTodosHandler)
	flag.BoolFunc("list-done", "list completed todo's.", cli.listDoneTodosHandler)
	flag.BoolFunc("list-all", "list all todo's.", cli.listAllTodosHandler)

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
