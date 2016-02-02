package main

import (
	"fmt"

	"github.com/aboodmm/todo"
	"github.com/fatih/color"
)

func main() {
	var s todo.Stack
	var item1 = todo.NewItem("hello, testing junk")
	var item2 = todo.NewItem("hello, testing shit")
	s.Push(item1)
	s.Push(item2)
	printStack(s)
}

func printStack(s todo.Stack) {
	var size = s.Size
	for i := 0; i < size; i++ {
		printItem(s.Top.Value)
		s.Pop()
	}
}

func printItem(i *todo.Item) {
	fmt.Println("Date:", i.DateAdded, "Message:", i.Message)
}
