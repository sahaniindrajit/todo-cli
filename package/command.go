package todo

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlag struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlag() *CmdFlag {
	cf := CmdFlag{}

	flag.StringVar(&cf.Add, "add", "", "Add a new Todo")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a Todo by index")
	flag.IntVar(&cf.Del, "del", -1, "specify a index to delete")
	flag.IntVar(&cf.Toggle, "toogle", -1, "specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all Todos")

	flag.Parse()

	return &cf
}

func (cf *CmdFlag) Execute(todos *Todos) {

	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: Invalid formate. use ID:Title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: Invalid index")
			os.Exit(1)
		}
		todos.Edit(index, parts[1])
	case cf.Toggle != -1:
		todos.Toogle(cf.Toggle)
	case cf.Del != -1:
		todos.Delete(cf.Del)
	default:
		fmt.Println("Invalid command")
	}
}
