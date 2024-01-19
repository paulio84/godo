package main

import (
	"fmt"
	"os"

	"github.com/paulio84/godo/internal/models/filter"
	"github.com/paulio84/godo/internal/service/db"
)

type cli struct {
	db      db.DBService
}

func NewCLI() *cli {
	return &cli{
		db: db.NewSQLiteService(),
	}
}

func main() {
	cli := NewCLI()
	cli.db.Connect()
	defer cli.db.Close()
}
