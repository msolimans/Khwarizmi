package removeduplicatesfromsortedlist

import "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *common.ListNode) *common.ListNode {
    //2 pointers 
    if head == nil {
        return head
    }
    
    current := head 
    next := head.Next 
    
    for next != nil {
        if current.Val == next.Val {
            current.Next = next.Next 
            next = next.Next 
        } else {
            current = next 
            next = next.Next 
        }
    }
    
    return head 
}