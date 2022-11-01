package OddEvenLinkedList

import (
	. "github.com/msolimans/khawarizmi/leetcode/common"
)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    podd := &ListNode{Val: 0} 
    peven := &ListNode{Val: 0}
    oddc := podd
    evenc := peven 
    
    c := &ListNode{Val: 0, Next: head} 
    i := 1
    
    for c.Next != nil {
        //etract 
        t := c.Next 
        //remove
        c.Next = c.Next.Next //0 - 2 3 
        //separate 
        t.Next = nil 
        
        if i % 2 == 0 {//even 
            evenc.Next = t
            evenc = evenc.Next
        } else {
            oddc.Next = t
            oddc = oddc.Next
        }  
        i++
    }
    
    oddc.Next = peven.Next 
    return podd.Next 
    
}