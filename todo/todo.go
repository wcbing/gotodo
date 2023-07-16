package todo

import (
	"errors"
	"fmt"
	"time"
)

type item struct {
	ID       int
	Done     bool
	Content  string
	Category string
	CreatAt  time.Time
	DoneAt   *time.Time
}

type Todos []item

// add a new todo
func (t *Todos) Add(content string, category string) {
	todo := item{
		ID:       len(*t) + 1,
		Done:     false,
		Content:  content,
		Category: category,
		CreatAt:  time.Now(),
		DoneAt:   nil,
	}
	*t = append(*t, todo)

}

// show all content of one todo found by id
func (t *Todos) ListbyID(id int) {
	a := (*t)[id]
	if a.Done {
		fmt.Printf("│%3d│☑️│%30s│%8s│%.10s│%.10s│\n", a.ID, a.Content, a.Category, a.CreatAt, a.DoneAt)
	} else {
		fmt.Printf("│%3d│⬜│%30s│%8s│%.10s│          │\n", a.ID, a.Content, a.Category, a.CreatAt)
	}
}

// delete a todo by id
func (t *Todos) Delete(id int) error {
	if id <= 0 || id > len(*t) {
		return errors.New("ID error")
	}
	for i := id; i < len(*t); i++ {
		(*t)[i].ID = i
	}
	*t = append((*t)[:id-1], (*t)[id:]...)

	return nil
}

// Update the todo's content and category
func (t *Todos) Update(id int, content string, category string) error {
	if id <= 0 || id > len(*t) {
		return errors.New("ID error")
	}
	(*t)[id-1].Content = content
	(*t)[id-1].Category = category

	return nil
}

func (t *Todos) Markup(id int, state bool) error {
	if id <= 0 || id > len(*t) {
		return errors.New("ID error")
	}
	(*t)[id-1].Done = state
	if state {
		now := time.Now()
		(*t)[id-1].DoneAt = &now
	} else {
		(*t)[id-1].DoneAt = nil
	}
	return nil
}
