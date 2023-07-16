package cmd

import (
	"fmt"
	"strconv"
	"wcbing/gotodo/todo"
)

func Markup(t *todo.Todos, args []string) {
	if len(args) == 0 {
		fmt.Printf("Missing parameter\n\n")
		Help()
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid parameter")
		return
	}
	if id <= 0 || id >= len(*t) {
		fmt.Print("ID error!")
		return
	}
	if len(args) == 1 || args[1] == "done" {
		t.Markup(id, true)
	} else if args[1] == "undone" {
		t.Markup(id, false)
	} else {
		fmt.Println("Invalid parameter")
		return
	}
}
