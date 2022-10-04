package pathsum

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func hasPathSum(root *TreeNode, targetSum int) bool {
    
    return hasPathSumCumulative(root, 0, targetSum)
    
}

func hasPathSumCumulative(root *TreeNode, sum, targetSum int) bool {
    if root == nil {
        return false
    }
    
    if  root != nil {
        root.Val = sum + root.Val 
    }
    
    //leaf is a node with no children 
    if root.Left == nil && root.Right == nil {//leaf 
        return root.Val == targetSum 
    }
    
    return hasPathSumCumulative(root.Left, root.Val, targetSum) || hasPathSumCumulative(root.Right, root.Val, targetSum)

}