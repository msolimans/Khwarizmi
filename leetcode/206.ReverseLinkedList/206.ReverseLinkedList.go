package reverselinkedlist

import (
	"container/list"

	. "github.com/msolimans/khawarizmi/leetcode/common"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	//use stack  
	stack := list.New()
	//lifo
	for head != nil {
		stack.PushBack(head.Val)
		head = head.Next
	}

	pres := &ListNode{}
	c := pres 
	for stack.Back() != nil {
		back := stack.Back()
		c.Next = &ListNode {Val: back.Value.(int)}
		c = c.Next
		stack.Remove(back)
	}

	return pres.Next
}

///inplace reversal 
func reverseList2(head *ListNode) *ListNode {
	//reverse using constant auxliary space 
	pre := &ListNode{} 

	c := head 
	for c != nil {
		//move current 
		//preserve next 
		t := c.Next
		c.Next = nil 
		
		//put all res array after c 
		c.Next = pre.Next 
		//put c in front 
		pre.Next = c 

		//update c 
		c = t 
	}

	return pre.Next

}