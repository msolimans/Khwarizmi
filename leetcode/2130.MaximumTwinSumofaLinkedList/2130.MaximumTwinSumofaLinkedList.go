package maximumtwinsumofalinkedlist

import (
	"container/list"

	. "github.com/msolimans/khawarizmi/leetcode/common"
)

func pairSum(head *ListNode) int {

    stack := list.New()
    fast, slow, c := head, head, head
    //get the middle 
	for fast != nil && fast.Next != nil {
      fast = fast.Next.Next
      slow = slow.Next  
    }
	//push second half of linked List in the stack 
    for slow != nil {
        stack.PushFront(slow)
		slow = slow.Next
    }
    
	//move consistently now from begining of linked list with the stack's pops 
	c = head
    max := c.Val
    for stack.Len() != 0 {
        front := stack.Front()
        twin := front.Value.(*ListNode)
        stack.Remove(front)

		//calc the max 
        if c.Val + twin.Val > max {
          max = c.Val + twin.Val
        }
		
		c = c.Next
    }

    return max
}

///////


//2nd solutions without extra space 
//using reverse second half of linked list 


func pairSum2(head *ListNode) int {

    preHead := &ListNode{Next: head }

    fast, slow := preHead,preHead
    //get the middle 
	for fast != nil && fast.Next != nil {
      fast = fast.Next.Next
      slow = slow.Next  
    }

	secondHalf := slow.Next
	slow.Next = nil 

	reversedHalf := reverse(secondHalf)
	slow.Next = reversedHalf
	
	first := preHead.Next
	slow = slow.Next
	max := first.Val

	for slow != nil {
		if slow.Val + first.Val > max { 
			max = slow.Val + first.Val
		}
		//move together
		first = first.Next
		slow = slow.Next
	}

	return max 
}




func reverse (node *ListNode) *ListNode {
	//now separate and reverse
	preHead := &ListNode{}
	c := node 

	for c!=nil {
		cn := c.Next
		c.Next = preHead.Next
		preHead.Next = c 
		c = cn 
	}

	return preHead.Next
}