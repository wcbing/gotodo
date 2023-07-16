package cmd

import (
	"fmt"
	"wcbing/gotodo/todo"
)

func Add(t *todo.Todos, args []string) {
	if len(args) == 0 {
		fmt.Printf("Missing parameter\n\n")
		Help()
		return
	}
	if len(args) == 1 {
		t.Add(args[0], "Default")
	} else {
		t.Add(args[0], args[1])
	}
	fmt.Println("Successful added!")
	printHead()
	printByIndex(t, len(*t)-1)
	printTail()
}
