package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAT *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAT: nil,
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("index out of range")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil

}

func (todos *Todos) toggle(index int) error {
	t := (*todos)
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAT = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := !t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAT = &completionTime
	}
	t[index].Title = title
	return nil

}

func (todos *Todos) print() {
	mainTable := table.New(os.Stdout)
	mainTable.SetRowLines(false)
	mainTable.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Completed {
			completed = "✅"
			if t.CompletedAT != nil {
				completedAt = t.CompletedAT.Format(time.RFC1123)
			}
		}

		mainTable.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	mainTable.Render()

	// incompleteTable := table.New(os.Stdout)
	// incompleteTable.SetRowLines(false)
	// incompleteTable.SetHeaders("Incomplete Todos")

	// incompleteTodos := 0
	// for _, t := range *todos {
	// 	if !t.Completed {
	// 		incompleteTodos++
	// 	}
	// }

	// if incompleteTodos > 0 {
	// 	incompleteTable.AddRow(fmt.Sprintf("%d todos not completed", incompleteTodos))
	// }

	// incompleteTable.Render()
}
