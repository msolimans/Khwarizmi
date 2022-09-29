package palindromelinkedlist

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	//go to middle - separate into 2 lists - reverse the second half of the linked list then compare both parts 

	preHead := &ListNode{Next: head}

	slow, fast := preHead, preHead
	//0,1,2,2,1
	//////^
	//0,1,2,3,2,1
	//

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//separate 
	secondList := slow.Next
	slow.Next = nil 

	reversed := reverse(secondList)
	slow = head 

	//now we start comparing (should be of the same size and element's value match)
	for slow != nil && reversed != nil {
		if slow.Val != reversed.Val {
			return false 
		}
		slow = slow.Next
		reversed = reversed.Next
	}

	
	//both of em should be pointing to nil (we may have identical lists but one's shorter than the other)
	//0,1,2,3,2,1 (only one case we need to check: odd numbers so we may have slow points to one "Next" elem, in such case we will return true)
	return (slow == nil && reversed == nil) || (slow != nil && slow.Next == nil && reversed == nil)
}

func reverse(head *ListNode) *ListNode {
	preHead := &ListNode{}

	c := head 
	for c != nil {
		temp := c.Next
		
		c.Next = preHead.Next
		preHead.Next = c 

		c = temp 
	}
	return preHead.Next 
}
