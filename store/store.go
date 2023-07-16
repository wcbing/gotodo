package store

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wcbing/gotodo/todo"
)

var homeDir, _ = os.UserHomeDir()
var fpath = filepath.Join(homeDir, ".config/gotodo/")
var file = filepath.Join(homeDir, ".config/gotodo/data.json")

func Store(t *todo.Todos) bool {
	data, err := json.Marshal(t)
	if err != nil {
		fmt.Print(err)
		return false
	}
	if _, err = os.Stat(fpath); os.IsNotExist(err) {
		if err = os.Mkdir(fpath, 0755); err != nil {
			fmt.Print(err)
			return false
		}
	}
	if _, err = os.Stat(file); os.IsNotExist(err) {
		if _, err = os.Create(file); err != nil {
			fmt.Print(err)
			return false
		}
	}
	if os.WriteFile(file, data, 0666) != nil {
		fmt.Print(err)
		return false
	}
	return true
}

func Restore(t *todo.Todos) bool {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err)
		fmt.Printf("\nWill create data file in %s\n", file)
		return false
	}
	err = json.Unmarshal(data, t)
	if err != nil {
		fmt.Print(err)
		return false
	}
	return true
}
