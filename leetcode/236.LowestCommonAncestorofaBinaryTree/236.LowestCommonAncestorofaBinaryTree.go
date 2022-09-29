package lowestcommonancestorofabinarytree

import (
	. "github.com/msolimans/khawarizmi/leetcode/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    //find target nodes first 
    if root == nil || root == p || root == q {
        return root 
    }
    
	//which sides target node exist (if one in the left and one in the right we return their parent, otherwise we check which is not nil)
    left := lowestCommonAncestor(root.Left, p ,q)
    right := lowestCommonAncestor(root.Right, p ,q)
    
    //once both found, propagate up till we get least ancestor  
    
    if left == nil {
        return right 
    } else if right == nil {
        return left 
    } else {
        return root 
    }

}