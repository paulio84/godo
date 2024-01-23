package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"github.com/paulio84/godo/internal/models/filter"
	"github.com/paulio84/godo/internal/service/database"
)

type cli struct {
	dbService database.DBServicer
	command   Commander
}

func newCLI(db *sql.DB) *cli {
	return &cli{
		dbService: database.NewSQLiteService(db),
	}
}

func main() {
	db, err := openDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	cli := newCLI(db)
	cli.dbService.Init()

// cli.dbService.AddTodo("Gem Gem")
// cli.dbService.ToggleCompleted(2)
// l, _ := cli.dbService.ListTodos(filter.All)
// fmt.Println(l)
// cli.dbService.EditTodo(1, "Gem Gem")
// cli.dbService.PurgeTodos()

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
