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
	if id <= 0 || id > len(*t) {
		fmt.Print("ID error!")
		return
	}
	content, category := (*t)[id-1].Content, (*t)[id-1].Category
	var in string
	fmt.Println("(Leave blank will not change)")
	fmt.Print("Please input the content:")
	fmt.Scanf("%s", &in)
	if in != "" {
		content = in
	}
	fmt.Print("Please input the categoty:")
	fmt.Scanf("%s", &in)
	if in != "" {
		category = in
	}
	t.Update(id, content, category)
}
