package main

import (
	"fmt"
	todo "todo-cli/package"
)

func main() {
	fmt.Println("Todo - CLI")

	todos := todo.Todos{}
	storage := todo.NewStaorage[todo.Todos]("todos.json")

	storage.Load(&todos)

	// todos.Add("aaa")
	// todos.Add("aaa2")
	// todos.Add("aaa3")

	// todos.Print()

	// todos.Delete(0)
	todos.Print()

	storage.Save(todos)

}
