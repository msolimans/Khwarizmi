package validatebinarysearchtree

import (
	"math"

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
 func isValidBST(root *TreeNode) bool {
    return isValid(root, math.MinInt, math.MaxInt)
}


func isValid(root *TreeNode, start, end int) bool {
    if root == nil {
        return true 
    }

    return root.Val > start && root.Val < end && isValid(root.Left, start, root.Val) && isValid(root.Right, root.Val, end)

}
//2 approaches here 
//1) level order traversal and add all elems in an array - check if the array is sorted 
//2) chk each elem and pass the range for each one (pre order traversal) and make sure each node is within that range otherwise return false 

