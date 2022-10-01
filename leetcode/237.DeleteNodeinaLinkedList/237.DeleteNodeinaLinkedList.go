package deletenodeinalinkedlist

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteNode(node *ListNode) {
    //assign value of node to next node's value 
    node.Val = node.Next.Val
    //remove next node from linkedList 
    node.Next = node.Next.Next 
    
}