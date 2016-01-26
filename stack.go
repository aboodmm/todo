package todo

import "time"

type Stack struct {
	maxSize     int
	stackArray  []Item
	top         int
	dateCreated time.Time
}

func NewStack(m int) *Stack {
	s := new(Stack)
	s.maxSize = m
	s.top = -1
	return s
}

func Push(i Item, s *Stack) {
	s.top += 1
	s.stackArray[s.top] = i
}

func Pop(s *Stack) []Item {
	return s.stackArray
}
