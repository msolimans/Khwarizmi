package reorderlist

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

 func reorderList(head *ListNode)  {
	 //stack to reverse the second half of linkedList 
	 //then start distributing items properly 
	 
	//  c := head 
	

	pmed, med := findMedium(head)
	//only 1 elem 
	if pmed == med && med == head {
		return 
	}

	//2 elems
	// if pmed == head && med == head.Next && head.Next != nil && head.Next.Next == nil {
	// 	return
	// }

	pmed.Next = nil 


	//reverse everything using after med 
	stack := list.New()
	for med != nil {
		t := med.Next
		med.Next = nil 
		stack.PushBack(med)
		med = t 
	}

	c := head
 
	for c != nil && c.Next != nil {
		insertAfter(c, stack.Back().Value.(*ListNode))
		stack.Remove(stack.Back())
		c = c.Next.Next //2 jumps 
	}

	//insert left over from stack 
	for stack.Len() > 0 {
		insertAfter(c, stack.Back().Value.(*ListNode))
		stack.Remove(stack.Back())
		c = c.Next //1 jump 
	}
	
 }

 func insertAfter(node, what *ListNode) {
	t := node.Next 
	node.Next = what 
	what.Next = t 
 }

 func findMedium(head *ListNode) (*ListNode, *ListNode) { 

	pslow, slow, fast := head, head, head 

	for fast != nil && fast.Next != nil {
		pslow = slow 
		slow = slow.Next 
		fast = fast.Next.Next 
	}

	return pslow, slow


 }