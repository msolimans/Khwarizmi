package builtin

import (
	"container/list"
)

type stack struct {
	list.List
}

func NewStack() *stack {
	return &stack{}
}

func (s *stack) Push(i any) {
	s.PushFront(i)
}


func (s *stack) Pop() (any, bool) {
	if s.IsEmpty() {
		return nil, false 
	}
	front := s.Front()
	s.Remove(front)
	return front.Value, true 
}

func (s *stack) Peek() (any, bool) {
	if  s.IsEmpty() {
		return nil, false
	}  
	return s.Front().Value, true 
}

func (s *stack) IsEmpty() bool { 
	return s.Len() == 0 
}