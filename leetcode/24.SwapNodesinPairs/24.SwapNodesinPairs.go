package swapnodesinpairs

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil }
	if head.Next == nil {
		return head 
	}

	preHead := &ListNode{Next: head}

	c := preHead

	for c.Next != nil && c.Next.Next != nil {
		swap(c)
		c = c.Next.Next
	}

	return preHead.Next
}


func swap(head *ListNode) {
	temp := head.Next 
	temp2 := head.Next.Next
	head.Next = temp2.Next
	
	//separate 
	temp.Next = nil
	temp2.Next = nil 

	tail := head.Next //rest of linked List 
	head.Next = temp2
	temp2.Next = temp
	temp.Next = tail 

}