package twosumivinputisabst

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
func findTarget(root *TreeNode, k int) bool {
    arr := []int{}
    
	//build the array of numbers (Inline order traversal generating it as sorted)
    buildTree(root, &arr)
    i,j :=0,len(arr)-1

	//we make BS to combine 2 numbers that sum up to target k 
    for i < j {
        if arr[i] + arr[j] == k{
            return true 
        } else if arr[i] + arr[j] < k {
            i++
        } else {
            j--
        }
    }
    
    return false 
    
}

func buildTree(root *TreeNode, arr *[]int) {
    
    if root == nil {
        return 
    }
    
    buildTree(root.Left, arr)
    *arr = append(*arr,root.Val)
    buildTree(root.Right, arr)
}
