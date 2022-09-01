package datastructure

import "container/list"

type LLQueue struct {
	list list.List
} 

func (l *LLQueue) Enqueue(s any) {
	l.list.PushBack(s)
}

func (l *LLQueue) Len() int { 
	return l.list.Len()
}

func (l *LLQueue) Dequeue() any {
	front := l.list.Front()
	res := front.Value
	l.list.Remove(front)
	return res
}



