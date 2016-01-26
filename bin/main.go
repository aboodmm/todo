package main

import "fmt"
import "github.com/aboodmm/todo"

func main() {
	s := newStack(5)
	fmt.Printf("HELLO %d", s.maxSize)
}
