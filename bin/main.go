package main

import (
	"fmt"

	"github.com/aboodmm/todo"
	//	"os"
)

func main() {
	var s todo.Stack
	//	var item1 = todo.NewItem("hello, testing junk")
	//	var item2 = todo.NewItem("hello, testing shit")
	for {
		var command string
		fmt.Println("enter command")
		fmt.Scanf("%s", &command)
		switch command {
		case "push":
			pushItem(&s)
		case "pop":
			popItem(&s)
		case "print":
			printStack(s)
		}
	}
}

func printStack(s todo.Stack) {
	var size = s.Size
	for i := 0; i < size; i++ {
		printItem(s.Top.Value)
		s.Pop()
	}
	return
}

func printItem(i *todo.Item) {
	fmt.Println("Date:", i.DateAdded, "Message:", i.Message)
}

func pushItem(s *todo.Stack) {
	fmt.Println("enter message to push")
	var i string
	fmt.Scanf("%s", &i)
	var pi = todo.NewItem(i)
	s.Push(pi)
	return
}

func popItem(s *todo.Stack) {
	s.Pop()
	fmt.Println("Stack popped")
	return
}
