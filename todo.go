package main

import "github.com/aboodmm/todo/stack"
import "github.com/aboodmm/todo/item"

func main() {
	s := Stack.newStack(5)
	fmt.Printf("HELLO %d", s.maxSize)
}
