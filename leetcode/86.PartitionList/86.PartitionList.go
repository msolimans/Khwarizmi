package partitionlist

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
 func partition(head *ListNode, x int) *ListNode {
    //2 queues (one for less and one for greater values)
    
	preHead := &ListNode{Next: head}
	c := preHead 
	preHigh := &ListNode{}
	hc := preHigh

	for c != nil && c.Next != nil {

		if c.Next.Val < x {
			c = c.Next
		} else {
			injectNodeAfter(hc, c.Next.Val)
			hc = hc.Next
			deleteNodeAfter(c)
		}
		
	}

	if c.Next == nil {
		c.Next = preHigh.Next
	} else {
		c.Next.Next = preHigh.Next
	}
    
    return preHead.Next
}


func deleteNodeAfter(n *ListNode) {
    t := n.Next
    n.Next = n.Next.Next 
    t.Next = nil 
}
func injectNodeAfter(n *ListNode, val int) {
    t := n.Next
    n.Next = &ListNode{Val: val, Next: t}
}