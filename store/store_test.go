package store

import (
	"testing"

	"github.com/wcbing/gotodo/cmd"
	"github.com/wcbing/gotodo/todo"
)

func TestStore(t *testing.T) {
	a := &todo.Todos{}
	a.Add("hello", "a")
	a.Add("demo", "a")
	a.Add("hjkl", "b")
	a.Add("qwer", "c")
	a.Markup(2, true)

	res := Store(a)
	if res {
		t.Logf("Save successful")
	} else {
		t.Fatalf("Save failed")
	}
}

func TestRestore(t *testing.T) {
	a := &todo.Todos{}
	if Restore(a) {
		t.Logf("Read successful")
		cmd.List(a, nil)
	} else {
		t.Fatalf("Read failed")
	}

}
