package todo

import "time"
import "fmt"

type Item struct {
	DateAdded time.Time
	Message   string
}

func NewItem(m string) *Item {
	i := new(Item)
	i.DateAdded = time.Now()
	i.Message = m
	return i
}

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
