package implementstackusingqueues

import "container/list"

//let's implement queue using linked list here
type Queue struct {
    list *list.List
}

func InitQueue() *Queue {
    return &Queue {list: list.New()}
}

func (q *Queue) Len() int {
	return q.list.Len()
}

func (q *Queue) Enqueue(val int) {
    q.list.PushBack(val)
}

func (q *Queue) Dequeue() int {
    if q.Len() > 0 {
		node := q.list.Front()
		res := node.Value.(int)
		q.list.Remove(node)
		return res 
	}

	return -1 
}

////////////////////////////////////////////////////////////////
//queue is FIFO (see implementation above)
//while stack is FILO or LIFO

type MyStack struct {
    x *Queue
	y *Queue
	top int 
}

func Constructor() MyStack {
	return MyStack{
		x: InitQueue(),
		y: InitQueue(),
		top: -1,
	}
}


func (this *MyStack) Push(x int)  {
    //enqueue into the current/active queue 
	q := this.x
	if this.y.Len() > 0 {
		q = this.y
	}
	//update top 
	this.top = x 
	q.Enqueue(x)
}



func (this *MyStack) Pop() int {
	if this.Empty() {//current stack is empty 
		return -1 
	}

    //for every pop, pop every elem from current/active queue (except the last one)
	q1 := this.x //q1 is the one that's not empty 
	q2 := this.y //this is the empty one 

	if this.y.Len() > 0 {//y is active 
		q1 = this.y 
		q2 = this.x
	}

	//only one left (top is now -1)
	if q1.Len() == 1 {
		this.top = -1 
	}

	//dequeue all elems from queue except the last one 
	for q1.Len() > 1 {
		last := q1.Dequeue()
		this.top = last //keep updating top until we reach out previous last element 
		q2.Enqueue(last)
	} 
	  
	//last elem 
	return q1.Dequeue()
}


func (this *MyStack) Top() int {
   return this.top 
}


func (this *MyStack) Empty() bool {
    return this.x.Len() == 0 && this.y.Len() == 0 
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */