package main

import (
	"fmt"
	"os"

	"github.com/wcbing/gotodo/cmd"
	"github.com/wcbing/gotodo/store"
	"github.com/wcbing/gotodo/todo"
)

func main() {

	var t = &todo.Todos{}
	store.Restore(t)

	if len(os.Args) == 1 {
		cmd.List(t, nil)
	} else {
		switch os.Args[1] {
		case "-h":
			fallthrough
		case "--help":
			fallthrough
		case "help":
			cmd.Help()
		case "-l":
			fallthrough
		case "list":
			cmd.List(t, os.Args[2:])
		case "-a":
			fallthrough
		case "add":
			cmd.Add(t, os.Args[2:])
		case "-u":
			fallthrough
		case "update":
			cmd.Update(t, os.Args[2:])
		case "-m":
			fallthrough
		case "markup":
			cmd.Markup(t, os.Args[2:])
		case "-d":
			fallthrough
		case "delete":
			cmd.Delete(t, os.Args[2:])
		default:
			fmt.Println("Invalid parameter")
			fmt.Println()
			cmd.Help()
		}
	}
	store.Store(t)
}
