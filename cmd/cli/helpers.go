package main

import (
	"fmt"

	"github.com/paulio84/godo/internal/models/core"
)

func displayHelpText() {
	fmt.Println("Some help text")
}

func displayTodoData(res core.DBResult) {
	if res.Err != nil {
		fmt.Println(res.Err)
	}

	// display data
	fmt.Println("RESULTS: ", res.Data)
}

func displayCreated(res core.DBResult) {
	if res.Err != nil {
		fmt.Println(res.Err)
	}

	fmt.Printf("Created %d todo item.\n", res.RowsAffected)
}
