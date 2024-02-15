package gotodo

import (
	"errors"
	"fmt"
)

type Todo struct {
	Title       string
	Description string
}

var todos = map[string]*Todo{}

// Create a todo
func createTodo(id int, todo *Todo) {
	todos[string(id)] = todo
}

// Read a todo
func readTodo(id int) (*Todo, error) {
	if todo := todos[string(id)]; todo == nil {
		return nil, errors.New("Read operation failed...")
	} else {
		return &Todo{Title: todo.Title, Description: todo.Description}, nil
	}
}

// Edit a todo
func editTodo(id int, t *Todo) {
	rtodo, err := readTodo(id)
	if err != nil {
		createTodo(id, t)
		return
	}
	rtodo.Title = t.Title
	rtodo.Description = t.Description

}

// Delete a todo
func deleteTodo(id int) {
	delete(todos, string(id))
}

func Run() {
	t := &Todo{Title: "Todo 1", Description: "A simple todo list"}
	createTodo(1, t)
	_, err := readTodo(1)
	if err != nil {
		fmt.Println(err)
	}
	editTodo(11, &Todo{Title: "Updated title", Description: "Updated Description"})
	for _, v := range todos {
		fmt.Println(v.Title, v.Description)
	}
	deleteTodo(11)
}
