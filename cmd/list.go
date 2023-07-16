package cmd

import (
	"fmt"

	"github.com/wcbing/gotodo/todo"
)

func List(t *todo.Todos, args []string) {
	printHead()
	if len(args) == 0 {
		for i := 0; i < len(*t); i++ {
			printByIndex(t, i)
		}
	} else {
		for index, v := range *t {
			if v.Category == args[0] {
				printByIndex(t, index)
			}
		}
	}
	printTail()
}

func printByIndex(t *todo.Todos, index int) {
	if index < 0 || index >= len(*t) {
		fmt.Print("ID error!")
		return
	}
	v := (*t)[index]
	if v.Done {
		fmt.Printf("│%3d│☑️│%30s│%8s│%.10s│%.10s│\n", v.ID, v.Content, v.Category, v.CreatAt, v.DoneAt)
	} else {
		fmt.Printf("│%3d│⬜│%30s│%8s│%.10s│          │\n", v.ID, v.Content, v.Category, v.CreatAt)
	}
}
func printHead() {
	fmt.Println("┌───┬──┬──────────────────────────────┬────────┬──────────┬──────────┐")
	fmt.Println("│ Id│⬜│                      Content │Category│Created At│   Done At│")
	fmt.Println("├───┼──┼──────────────────────────────┼────────┼──────────┼──────────┤")

}
func printTail() {
	fmt.Println("└───┴──┴──────────────────────────────┴────────┴──────────┴──────────┘")

}
