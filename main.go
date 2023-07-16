package main

import (
	"fmt"
	"os"
	"wcbing/gotodo/cmd"
	"wcbing/gotodo/todo"
)

func main() {

	//test data
	var t = &todo.Todos{}
	t.Add("hello", "a")
	t.Add("demo", "a")
	t.Add("jdfhjskdfhj", "b")
	t.Markup(2, true)

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

}
