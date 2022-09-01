package common

type ListNode struct {
	Val int 
	Next *ListNode 
}

// The following functions are used to simplify unit tests

func Init(ins []int) *ListNode {
	if len(ins) == 0 {
		return nil 
	}

	preHead := &ListNode{}
	c := preHead
	for _,v:= range ins {
		c.Next = &ListNode{Val: v}
		c = c.Next
	}
	return preHead.Next	
}

func (l *ListNode) Equals(r *ListNode) bool {
	for l != nil && r != nil {
		if l.Val != r.Val {
			return false 
		}
		l = l.Next 
		r = r.Next
	}

	return  l == nil && r == nil 
}
