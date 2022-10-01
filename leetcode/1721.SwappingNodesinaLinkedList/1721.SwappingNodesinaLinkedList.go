package swappingnodesinalinkedlist

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapNodes(head *ListNode, k int) *ListNode {
	//use prehead to get elements correctly 
    preHead := &ListNode{Val: 0, Next: head}
    
    fast, slow := preHead, preHead

	for i:= 0;i<k;i++ { 
        fast = fast.Next 
    }

	first := fast //now at k 

	for fast != nil { //move fast and slow together until fast = nil (slow should be at n-k)
		slow = slow.Next 
		fast = fast.Next
	}

	//swap values 
	t := slow.Val 
	slow.Val = first.Val
	first.Val = t

	return preHead.Next
    
}