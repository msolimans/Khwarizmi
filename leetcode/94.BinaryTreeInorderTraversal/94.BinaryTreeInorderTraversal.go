package binarytreeinordertraversal

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for a binary tree node. check github.com/msolimans/khawarizmi/leetcode/common
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var n = []int{}
func inorderTraversal(root *TreeNode) []int {
    //left side then node then right side 
	//fil in left side 
	
	inorder(root)
	return n

}

func inorder(c *TreeNode) {
	if c == nil {
		return
	}

	if c.Left != nil { 
		inorder(c.Left)
	}

	n = append(n, c.Val)

	if c.Right != nil {
		inorder(c.Right)
	}
}