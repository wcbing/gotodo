package main

import (
	"os"
	"wcbing/gotodo/cmd"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "help" {
			cmd.Help()
		}
	}

}
