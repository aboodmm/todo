package main

import "fmt"
import "github.com/aboodmm/todo"

func main() {
	s := todo.NewStack(5)
	fmt.Printf("Made new stack with size %d", s.MaxSize)
	testItem := todo.NewItem("fuck yeah")
	todo.Push(testItem, s)
}
