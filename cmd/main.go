package main

import (
	"fmt"
	todo "todo-cli/package"
)

func main() {
	fmt.Println("Todo - CLI")

	todos := todo.Todos{}

	todos.Add("aaa")
	todos.Add("aaa2")
	todos.Add("aaa3")

	todos.Print()

	todos.Delete(0)
	todos.Print()

}
