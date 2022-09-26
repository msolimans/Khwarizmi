package linkedlistcycleii

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func detectCycle(head *ListNode) *ListNode {
    //slow, fast and previous fast pointers 
   // var pfast *ListNode = nil 
    slow, fast := head, head 
    hasCycle := false 
    
    for fast != nil && fast.Next != nil {
        
        slow = slow.Next 
    //    pfast = fast 
        fast = fast.Next.Next 
        
        if slow == fast {
            hasCycle = true 
            break 
        }
    }
    
    
    if !hasCycle {
        return nil 
    }
    
    //now slow equals fast - move normally but start from head 
    
    slow = head;
     
    for slow != fast {
        slow=slow.Next
        fast=fast.Next
    }
    
    return slow

}


