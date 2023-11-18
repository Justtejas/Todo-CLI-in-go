package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/justtejas/Todo-CLI-in-go"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new task")
	flag.Parse()

	todos := &todo.Todo{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch{
	case *add:
		todos.Add("Just a test")
		err := todos.Store(todoFile)
		if err != nil {
			fmt.Println("Error")
			os.Exit(1)
		} 
	default:
		fmt.Println("Invalid Command")
		os.Exit(0)
	}
}