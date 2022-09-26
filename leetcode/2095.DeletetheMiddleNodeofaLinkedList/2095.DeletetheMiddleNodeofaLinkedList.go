package deletethemiddlenodeofalinkedlist

import . "github.com/msolimans/khawarizmi/leetcode/common"


func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil 
    }
    
    pslow, slow, fast := head, head, head 
    
    for fast != nil && fast.Next != nil {
        pslow = slow 
        slow = slow.Next 
        fast = fast.Next.Next 
    }
    
    //previous slow: delete next node 
    pslow.Next = slow.Next
    //medium 
    slow.Next = nil 
    return head 
}
