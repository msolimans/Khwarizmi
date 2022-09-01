package datastructure

type LinkedList struct {
	root *LinkedListNode
}

type LinkedListNode struct {
	val int 
	next *LinkedListNode
}

func Init() *LinkedList {
	return &LinkedList{
		root: nil,
	}
}

func (l *LinkedList) PushBack(i int) {
	if l.root == nil {
		l.root = &LinkedListNode{val: i}
		return 
	}
	
	c := l.root
	for c != nil && c.next != nil {
		c = c.next
	}

	c.next = &LinkedListNode{val: i }
}

func (l *LinkedList) PushFront(i int ) {
	c := &LinkedListNode{val: i}
	c.next = l.root
	l.root = c 
}

func (l *LinkedList) PopFront() (int, bool) {
	if l.root == nil {
		return -1, false 
	}

	c := l.root
	l.root = l.root.next
	return c.val, true
}

func (l *LinkedList) PopLast() (int, bool) {
	if l.root == nil {
		return -1, false 
	}

	pc := l.root 
	c := l.root.next

	for c != nil && c.next != nil {
		pc = c 
		c = c.next
	}
	
	if (pc != nil && c == nil) {//only 1 elem 
		res := l.root.val
		l.root = nil 
		return res, true 
	}
	
	pc.next = nil 
	return c.val, true 

}