package designcircularqueue

type MyCircularNode struct {
    Value int 
    Next *MyCircularNode
}
type MyCircularQueue struct {
    Root *MyCircularNode    
    Current *MyCircularNode
    Size int 
    Length int 
}


func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{Root: nil, Current: nil, Size: k, Length: k}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.Size == 0 {
        return false
    }
    
    if this.Root == nil {
        this.Root = &MyCircularNode{Value: value}
        this.Current = this.Root

    } else {
        this.Current.Next = &MyCircularNode{Value: value, Next: this.Root}
        this.Current = this.Current.Next 
    }
    
    this.Size--
    return true 
}


func (this *MyCircularQueue) DeQueue() bool {
    if this.IsEmpty() {
        return false 
    }
    
    if this.Root == nil {
        return false 
    }
    
    if this.Root == this.Current {
        this.Root = nil 
        this.Current = nil 
        this.Size++
        return true 
    }
    
    this.Root = this.Root.Next 
    this.Size++
    return true 
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty(){
		return -1 
	}
    return this.Root.Value 
}


func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty(){
		return -1 
	}
	return this.Current.Value
}


func (this *MyCircularQueue) IsEmpty() bool {
    return this.Size == this.Length 
}


func (this *MyCircularQueue) IsFull() bool {
    return this.Size == 0
}


/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */