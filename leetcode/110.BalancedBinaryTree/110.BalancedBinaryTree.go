package balancedbinarytree

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

var m = map[*TreeNode]bool{}
//dynamic programming
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true 
	}	 
	
	//post order traversal (this can be enhanced by dynamic programming by preserving hashMap of treeNode and bool)
	left := false 
	if _,exists := m[root.Left]; exists {
		left = m[root.Left]
	} else {
		left = isBalanced(root.Left)
	}
	

	right := false 
	if _,exists := m[root.Right]; exists {
		right = m[root.Right]
	} else {
		right = isBalanced(root.Right)
	}

	if math.Abs(height(root.Left) - height(root.Right)) > 1 {
		return false 
	}
	
	return left && right 
    
}

 


func height(node *TreeNode) float64 {
	if node == nil {
		return 0 
	}

	return 1 + math.Max(height(node.Left), height(node.Right))
}