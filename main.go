package main

import (
	"os"
	"wcbing/gotodo/cmd"
	"wcbing/gotodo/todo"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help" {
			cmd.Help()
		}
	}

	var a todo.Todos
	a.Add("hello", "a")
	a.Add("jdfhjskdfhj", "b")
	a.Markup(1, true)

	for i := 0; i < len(a); i++ {
		a.ListbyID(i)
	}

}
