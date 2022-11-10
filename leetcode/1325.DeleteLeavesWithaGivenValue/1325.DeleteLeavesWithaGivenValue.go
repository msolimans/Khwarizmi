package deleteleaveswithagivenvalue

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    if root == nil {
       return nil 
    }
    //post order traversal    
    left := removeLeafNodes(root.Left, target)
    right := removeLeafNodes(root.Right, target)

    //assign new left and right to current root 
    root.Left = left 
    root.Right = right 
   
    //case of removing node is we don't left or right (leaf node) and node's value = target 
   if root.Left == nil && root.Right == nil && root.Val == target {
       return nil 
   }
   //return root 
   return root 
}
