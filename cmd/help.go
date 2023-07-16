package cmd

import (
	"fmt"
)

func Help() {
	fmt.Println("gotodo, a simple ToDo software written by Go")
	fmt.Println("Go ToDo, GoTo Do")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("    -l, list [category]")
	fmt.Println("        List all todos (in this category)")
	fmt.Println("    -a, add <todo content> [category]")
	fmt.Println("        Add a new todo (in this category)")
	fmt.Println("    -u, update <id>")
	fmt.Println("        Update an existing todo")
	fmt.Println("    -m, markup <id> [undone]")
	fmt.Println("        Mark as done (or undone)")
	fmt.Println("    -d, delete <id>/all")
	fmt.Println("        Delete an/all existing todo")
	fmt.Println("    -h, --help, help")
	fmt.Println("        Show this help message")
	fmt.Println("")
}
