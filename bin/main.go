package main

import "fmt"
import "github.com/aboodmm/todo"

func main() {
	s := todo.NewStack(5)
	fmt.Printf("Made new stack with size %d\n", s.MaxSize)
	var testItem = todo.NewItem("fuck yeah")
	fmt.Printf(testItem.Message)
	s.Push(testItem)
}
