package todo

import (
	"errors"
	"fmt"
	"time"
)

type item struct {
	tid      int
	ID       int
	Done     bool
	Content  string
	Category string
	CreatAt  time.Time
	DoneAt   *time.Time
}

type Todos []item

var num int
var nextTid int

// add a new todo
func (t *Todos) Add(content string, category string) {
	num++
	todo := item{
		tid:      nextTid,
		ID:       num,
		Done:     false,
		Content:  content,
		Category: category,
		CreatAt:  time.Now(),
		DoneAt:   nil,
	}
	*t = append(*t, todo)

	nextTid++
}

// show all content of one todo found by id
func (t *Todos) ListbyID(id int) {
	a := (*t)[id]
	if a.Done {
		fmt.Printf("|%3d|☑️|%30s|%9s|%.10s|%.10s|\n", a.ID, a.Content, a.Category, a.CreatAt, a.DoneAt)
	} else {
		fmt.Printf("|%3d|⬜|%30s|%9s|%.10s|          |\n", a.ID, a.Content, a.Category, a.CreatAt)
	}
}

// delete a todo by id
func (t *Todos) Delete(id int) error {
	if id <= 0 || id > num {
		return errors.New("ID error")
	}
	for i := id; i < num; i++ {
		(*t)[i].ID = i - 1
	}
	*t = append((*t)[:id-1], (*t)[id:]...)

	return nil
}

// Update the todo's content and category
func (t *Todos) Update(id int, content string, category string) error {
	if id <= 0 || id > num {
		return errors.New("ID error")
	}
	(*t)[id].Content = content
	(*t)[id].Category = category

	return nil
}

func (t *Todos) Markup(id int, state bool) error {
	if id <= 0 || id > num {
		return errors.New("ID error")
	}
	(*t)[id].Done = state
	if state {
		now := time.Now()
		(*t)[id].DoneAt = &now
	} else {
		(*t)[id].DoneAt = nil
	}
	return nil
}
