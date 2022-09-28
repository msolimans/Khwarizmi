package removelinkedlistelements

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	
	preHead := &ListNode{Next: head } 

	c := preHead

	for c.Next != nil {
		if c.Next.Val == val {
			//remove that node that has value of "val"
			c.Next = c.Next.Next
			continue
			//don't move we need to see the next as it might be val too
		}
		//move to next 
		c = c.Next
	}

	return preHead.Next
}
