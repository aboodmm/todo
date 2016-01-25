package main

import (
	"github.com/aboodmm/todo/item"
	"github.com/aboodmm/todo/stack"
)

func main() {
	s := Stack.newStack(5)
	fmt.Printf("HELLO %d", s.maxSize)
}
