package stack

import "time"

type Stack struct {
	maxSize    int
	stackArray []struct{}
	top        int
}

func newStack(m int) *Stack {
	s := new(Stack)
	s.maxSize = m
	s.top = -1
	return s
}

func Push(n string, s *Stack) {
	s.top += 1
	s.stackArray[s.top] = n
}

func Pop(s *Stack) []struct{} {
	return s.stackArray
}
