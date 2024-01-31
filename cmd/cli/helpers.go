package main

import (
	"fmt"

	"github.com/paulio84/godo/internal/models/core"
)

func displayTodoData(res core.DBResult) {
	if res.Err != nil {
		fmt.Println(res.Err)
		return
	}

	// display data
	fmt.Println("RESULTS: ", res.Data)
}

func displayCreated(res core.DBResult) {
	if res.Err != nil {
		fmt.Println(res.Err)
		return
	}

	fmt.Printf("Created %d todo item.\n", res.RowsAffected)
}

func displayEdited(res core.DBResult) {
	if res.Err != nil {
		fmt.Println(res.Err)
		return
	}

	fmt.Printf("Updated %d todo item.\n", res.RowsAffected)
}

func displayPurged(res core.DBResult) {
	if res.Err != nil {
		fmt.Println(res.Err)
		return
	}

	fmt.Printf("Purged %d todo item(s).\n", res.RowsAffected)
}

func displayToggled(res core.DBResult) {
	if res.Err != nil {
		fmt.Println(res.Err)
		return
	}

	fmt.Printf("Toggled %d todo item(s).\n", res.RowsAffected)
}
