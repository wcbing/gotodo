package cmd

import (
	"fmt"
	"strconv"

	"github.com/wcbing/gotodo/todo"
)

func Delete(t *todo.Todos, args []string) {
	if len(args) == 0 {
		fmt.Printf("Missing parameter\n\n")
		Help()
		return
	}
	// delete all todos
	if args[0] == "all" {
		fmt.Println("Will delete all todos")
		fmt.Print("Are you sure?: (Y/n)")
		var a string
		fmt.Scanf("%s", &a)
		if a == "" || a == "Y" || a == "y" {
			*t = nil
		} else {
			fmt.Println("User cancels")
		}
		return
	}
	// delete by id
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Print("Invalid parameter")
		return
	}
	if id <= 0 || id > len(*t) {
		fmt.Print("ID error!")
		return
	}

	fmt.Println("Will delete:")
	printHead()
	printByIndex(t, id-1)
	printTail()
	fmt.Print("Are you sure?: (Y/n)")
	var a string
	fmt.Scanf("%s", &a)
	if a == "" || a == "Y" || a == "y" {
		t.Delete(id)
	} else {
		fmt.Println("User cancels")
	}
}
