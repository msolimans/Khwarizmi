package reverselinkedlistii

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	//one node only 
	if head == nil {
		return nil 
	}
	if head.Next == nil {
		return head 
	}

	//no need to do anything here 
	if left == right {
		return head 
	}

	//find boundary 
	innerLength := (right - left) + 1 // (cause we use preHead to adjust indexes and avoid nil exceptions)
	
	preHead := &ListNode{Val:0} //to solve edges anf have proper indexes 
	preHead.Next = head 

	c := preHead

	for left > 1 { //2 1 0,[1,2,3,4,5]
		c = c.Next  
		left-- 
	}

	//pf an pr are previous left/right nodes 
	prevLeftNode := c

	for innerLength > 0 {
		c = c.Next 
		innerLength-- 
	}
	
	rightNode := c 

	//separating out list 
	rhead := prevLeftNode.Next 
	prevLeftNode.Next = nil 

	tail := rightNode.Next
	rightNode.Next = nil 
	//////////////////////////////

	reversedH, reversedT := reverse(rhead)

	//connect them again 
	prevLeftNode.Next = reversedH
	reversedT.Next = tail 
	
	return preHead.Next 
}

//returns head and tail of reversed list 
func reverse(head *ListNode) (*ListNode, *ListNode) {
	pre := &ListNode{} 
	c := head 
	tail := head 

	for c!= nil {
		t := c.Next
		c.Next = nil  //separate curent node from next nodes 
		
		c.Next = pre.Next //current's next node points to pre.Next 
		pre.Next = c //change pre's next to current 

		c = t //move current to original next 
	}

	return pre.Next, tail 
}