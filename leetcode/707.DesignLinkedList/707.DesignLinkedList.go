package designlinkedlist

type MyLinkedListNode struct {
    Val int 
    Next *MyLinkedListNode 
}

type MyLinkedList struct {
	head *MyLinkedListNode
}


func Constructor() MyLinkedList {
    return MyLinkedList{}
}


func (this *MyLinkedList) Get(index int) int {
	//preHead := &MyLinkedListNode{Next: this.head}
	//1 -2 -3 
	//c - c		1 

	c := this.head
	for index > 0 {
		//we exceeded the list length so return -1 
        if c == nil || c.Next == nil {
            return -1
        }
		c = c.Next 
		index -- 
	}

	if c == nil {
		return -1 
	}

	return c.Val
}


func (this *MyLinkedList) AddAtHead(val int)  {
	newHead := &MyLinkedListNode{Val: val, Next: this.head}
	this.head = newHead
}


func (this *MyLinkedList) AddAtTail(val int)  {
	//we ca

	preHead := &MyLinkedListNode{Next: this.head}
	c := preHead
	//previous current since current will be pointing to nil at the end of the following loop 
	pc := preHead

	for c != nil {
		pc = c 
		c = c.Next
	}

	pc.Next = &MyLinkedListNode{Val: val}
	//we may have zero node in the list 
	this.head = preHead.Next
}


func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	//we add preHead to adjust indexes to start from zero 
    preHead := &MyLinkedListNode{Next: this.head}

	c := preHead

	//0 - 1 - 2 - 3  
	//c   c				 	1

	for index > 0 {
        if c.Next == nil {//if it passes lenght of linkedList, return without addition 
            return 
        }
		c = c.Next
		index -- 
	}

	temp := c.Next 
	c.Next = &MyLinkedListNode{Val: val, Next: temp}
	
	//we need to update head as we may have added a new node at index "0" - thus head should be updated with the new node 
	this.head = preHead.Next
}


func (this *MyLinkedList) DeleteAtIndex(index int)  {
	
	preHead := &MyLinkedListNode{Next: this.head}

	c := preHead
 
	//0 - 1 - 2 - 3 
	//    c				1
	//    c   c			2 

	for index > 0 { 
		//check to make sure we can delete that node at this specific index
		if c.Next == nil {
			return
		}
		c = c.Next
		index -- 
	}

	//check to make sure we can delete that node at this specific index
	if c.Next == nil {
		return
	}
	c.Next = c.Next.Next
	//we have to update head with preHead.Next since we may have deleted `head` itself 
	this.head = preHead.Next
}

