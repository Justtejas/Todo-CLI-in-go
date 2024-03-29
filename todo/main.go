package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	todo "github.com/justtejas/Todo-CLI-in-go"
)

const (
	todoFile = ".todos.json"
)

func main() {
	add := flag.Bool("add", false, "add a new task")
	complete := flag.Int("complete", 0, "complete a task by its number")
	del := flag.Int("del", 0, "delete a task by its number")
	list := flag.Bool("list", false, "list all tasks")
	flag.Parse()

	todos := &todo.Todo{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch{
		case *add:
			task, err := getInput(os.Stdin, flag.Args()...)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			todos.Add(task)
			err = todos.Store(todoFile)
			if err != nil {
				fmt.Println("Error")
				os.Exit(1)
			}
		case *complete > 0:
			if err := todos.Complete(*complete); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			err := todos.Store(todoFile)
			if err != nil {
				fmt.Println("Error")
				os.Exit(1)
			}
		case *del > 0:
			if err := todos.Delete(*del); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
			err := todos.Store(todoFile)
			if err != nil {
				fmt.Println("Error")
				os.Exit(1)
			}
		case *list:
			todos.Print()
			os.Exit(0)
		default:
			fmt.Println("Invalid Command")
			os.Exit(0)
	}
}

func getInput(r io.Reader, args ...string) (string, error){
	if len(args) > 0 {
		return strings.Join(args, " "), nil 
	}
	buf := bufio.NewScanner(r)
	buf.Scan()
	if err := buf.Err(); err != nil {
		return "", err
	}
	text := buf.Text()
	if len(text) == 0 {
		return "", fmt.Errorf("the task cannot be blank")
	}
	return text, nil
}