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
	cmgFlags := todo.NewCmdFlag()
	cmgFlags.Execute(&todos)
	storage.Save(todos)

}
