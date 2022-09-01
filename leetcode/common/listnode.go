package common

type ListNode struct {
	Val int 
	Next *ListNode 
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