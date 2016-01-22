package todo

import "time"

type Stack struct {
	stackArray [maxsize]struct{}
	maxSize    int
	top        int
}

func newStack(m int) Stack {
	s := new(Stack)
	s.maxSize = m
	s.top = -1
	return s
}

func Push(n string) {
	top += 1
	stackArray[top] = n
}

func Pop() struct{} {
	return stackArray
}
