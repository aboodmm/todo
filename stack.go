package todo

import "time"
import "fmt"

type Stack struct {
	Size int
	Top  *Item
}

func NewStack(m int) *Stack {
	s := new(Stack)
	return s
}

func (s *Stack) Push(i *Item) {
	s.Top = i
}

func (s *Stack) Pop() []*Item {
	return s.StackArray
}
