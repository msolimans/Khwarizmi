package datastructure

import (
	"container/list"
	"errors"
)

type LLStack struct {
	list *list.List
}

func InitLLStack() *LLStack {
	return &LLStack{list: list.New()}
}

func (l *LLStack) Push(val any) {
	l.list.PushFront(val)
}

func (l *LLStack) Pop() (any,error) {
	if l.list.Len() == 0{
		return nil, errors.New("Empty")
	}
	f := l.list.Front()
	l.list.Remove(f)
	return f.Value, nil

}

func (l *LLStack) IsEmpty() bool {
	return l.list.Len() == 0
}
