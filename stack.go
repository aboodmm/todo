package todo

import "time"

type Stack struct {
	MaxSize     int
	StackArray  []Item
	Top         int
	DateCreated time.Time
}

func NewStack(m int) *Stack {
	s := new(Stack)
	s.MaxSize = m
	s.Top = -1
	return s
}

func Push(i Item, s *Stack) {
	s.Top += 1
	s.StackArray[s.Top] = i
}

func Pop(s *Stack) []Item {
	return s.StackArray
}
