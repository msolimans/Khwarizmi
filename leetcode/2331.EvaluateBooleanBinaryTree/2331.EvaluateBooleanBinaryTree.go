package evaluatebooleanbinarytree

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
	//2 means OR 
	//3 means AND 
	//0 means false 
	//1 means true 
	if evalNode(root) == 0 {
		return false 
	}

	return true 
}

func evalNode(root *TreeNode) int {
	//handling children of leaf nodes 
	if root == nil {
		return -1 //to make sure this converts to true  
	}
	
	left := evalNode(root.Left)
	right := evalNode(root.Right)
	
	//handling leaf nodes themselves 
	if left == -1 && right == -1 {
		return root.Val
	}

	operand := root.Val
	root.Val = eval(left,right, operand)

	return root.Val
}

//should return either 0 or 1 
func eval(operator1, operator2, operand int) int {
	op1 := false 
	if operator1 == 1 {
		op1 = true 
	}

	op2 := false 
	if operator2 == 1 {
		op2 = true 
	}

	if operand == 2 {
		if op1 || op2 {
			return 1
		} else {
			return 0
		}
	} else {
		if op1 && op2 {
			return 1
		} else {
			return 0 
		}
	}
	

}