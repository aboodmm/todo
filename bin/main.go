package main

import "fmt"
import "github.com/aboodmm/todo"

func main() {
	s := todo.NewStack(5)
	fmt.Printf("HELLO %d\n", s.MaxSize)
}
