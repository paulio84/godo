package database

import (
	"database/sql"
	"errors"

	"github.com/paulio84/godo/internal/models/filter"
	"github.com/paulio84/godo/internal/models/todo"
)

type SQLiteService struct {
	db *sql.DB
}

func NewSQLiteService(db *sql.DB) DBServicer {
	return &SQLiteService{
		db: db,
	}
}

func (sqlite *SQLiteService) Init() error {
	if err := sqlite.executeTransaction(`
		CREATE TABLE IF NOT EXISTS todo (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title VARCHAR(100) NOT NULL,
			isCompleted INTEGER NOT NULL DEFAULT 0 CHECK(isCompleted IN (0, 1))
		);
	`); err != nil {
		return err
	}

	return nil
}

func (sqlite SQLiteService) AddTodo(title string) error {
	err := sqlite.executeTransaction("INSERT INTO todo (title) VALUES (?)", title)
	if err != nil {
		return err
	}

	return nil
}

func (sqlite SQLiteService) EditTodo(id int, title string) error {
	err := sqlite.executeTransaction("UPDATE todo SET title = ? WHERE id = ?", title, id)
	if err != nil {
		return err
	}

	return nil
}

func (sqlite SQLiteService) ToggleCompleted(id int) error {
	err := sqlite.executeTransaction("UPDATE todo SET isCompleted = ((isCompleted | 1) - (isCompleted & 1)) WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (sqlite SQLiteService) ListTodos(todoFilter filter.TodoFilter) ([]todo.Todo, error) {
	var query string
	switch todoFilter {
	case filter.All:
		query = "SELECT id, title, isCompleted FROM todo"
	case filter.OnlyCompleted:
		query = "SELECT id, title, isCompleted FROM todo WHERE isCompleted = 1"
	case filter.NotCompleted:
		query = "SELECT id, title, isCompleted FROM todo WHERE isCompleted = 0"
	}

	if query == "" {
		return nil, errors.New("cannot get list of todos")
	}

	rows, err := sqlite.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []todo.Todo
	for rows.Next() {
		var id int
		var title string
		var isCompleted int // use int here because SQLite doesn't support bool

		err = rows.Scan(&id, &title, &isCompleted)
		if err != nil {
			return nil, err
		}

		// append a new todo to the list
		todo := todo.Todo{
			ID:          id,
			Title:       title,
			IsCompleted: isCompleted != 0,
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (sqlite SQLiteService) PurgeTodos() error {
	err := sqlite.executeTransaction("DELETE FROM todo WHERE isCompleted = 1")
	if err != nil {
		return err
	}

	return nil
}

func (sqlite SQLiteService) executeTransaction(query string, args ...any) error {
	tx, err := sqlite.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}