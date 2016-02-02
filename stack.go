package todo

import ()

type Stack struct {
	Top  *Element
	Size int
}

type Element struct {
	Value *Item
	Next  *Element
}

func (s *Stack) Len() int {
	return s.Size
}

func (s *Stack) Push(Value *Item) {
	s.Top = &Element{Value, s.Top}
	s.Size++
}

func (s *Stack) Pop() (Value *Item) {
	if s.Size > 0 {
		Value = s.Top.Value
		s.Top = s.Top.Next
		s.Size--
		return
	}
	return nil
}
