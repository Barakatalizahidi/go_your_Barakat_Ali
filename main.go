package main

import "fmt"

type todo struct {
	id      uint
	title   string
	content string
}

// todos represents storage for todo
var todos []todo = []todo{}

func main() {
	addTodo(Todo{Id: }  ) 
}

// addTodo adds todo to the storage
func addTodo(newTodo todo) {
	//TODO: complete this function
	todos = append(todos, newTodo)
	fmt.Printf("Todo with ID %d added ", newTodo.id)
}

// getTodos displays all todo
func getTodos() {
	// TODO: complete this function

	fmt.Println("All Todos:")
	for _, todo := range todos {
		fmt.Printf("id: %d, title: %s, content: %s", todo.id, todo.title, todo.content)
	}
}

// getTodos displays all todo
func getTodoByID(i uint) {
	// TODO: complete this function

	for _, todo := range todos {
		if todo.id == i {
			fmt.Printf("id: %d, title: %s, content: %s", todo.id, todo.title, todo.content)
			return

		}
	}

	fmt.Printf("Todo with ID %d not found", i)
}

// deleteTodoByID deletes todo by given ID
func deleteTodoByID(i uint) {
	// TODO: complete this function

	for index, todo := range todos {
		for todo.id == i {
			todos = append(todos[:index], todos[index+1:]...)
			fmt.Printf("Todo with ID %d deleted", i)
			return

		}
	}
	fmt.Printf("Todo with ID %d not found", i)

}
