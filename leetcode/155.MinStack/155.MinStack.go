package minstack

//implement it in a linkedlist style
type MinNode struct {
	Val int 
	Next *MinNode
	MinVal int //at time of push what was the min value 	
}

type MinStack struct {	
	head *MinNode 
}


func Constructor() MinStack {
	return MinStack { head: nil}
}


func (this *MinStack) Push(val int)  {
    if this.head == nil {
		this.head = &MinNode{Val: val, MinVal: val}
		return
	}

	//time to add one more node (look at current min before push and insert the least minimum one here )
	node := &MinNode {
		Val: val, 
		MinVal: val,
	}

	if this.head.MinVal < val {
		node.MinVal = this.head.MinVal
	}  

	node.Next = this.head 
	this.head = node 
}


func (this *MinStack) Pop()  {
    //update head 	
	if this.head != nil {
		this.head = this.head.Next
	}
}


func (this *MinStack) Top() int {
    if this.head == nil {
		return -1
	}
	return this.head.Val
}


func (this *MinStack) GetMin() int {

	if this.head == nil {
		return -1
	}
	return this.head.MinVal   
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */