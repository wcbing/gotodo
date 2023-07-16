package cmd

import (
	"fmt"
	"strconv"

	"github.com/wcbing/gotodo/todo"
)

func Update(t *todo.Todos, args []string) {
	if len(args) == 0 {
		fmt.Printf("Missing parameter\n\n")
		Help()
		return
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Print("Invalid parameter")
		return
	}
	if id <= 0 || id >= len(*t) {
		fmt.Print("ID error!")
		return
	}
	var content, category string
	fmt.Println("(Leave blank will not change)")
	fmt.Print("Please input the content:")
	fmt.Scanf("%s", &content)
	fmt.Print("Please input the categoty:")
	fmt.Scanf("%s", &category)
	t.Update(id, content, category)
}
