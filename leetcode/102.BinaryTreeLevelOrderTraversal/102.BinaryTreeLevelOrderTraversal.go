package binarytreelevelordertraversal

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    //level by level needs BFS traversal 
	//use queue for this but it doesn't help in getting boundaries of tree 

	if root == nil {
		return [][]int{}
	}
	
	resultList := [][]int{}
	levelList := []*TreeNode{} 

	levelList = append(levelList, root)

	for {
		res, next := getLevelNodes(levelList)
		resultList = append(resultList, res)
		if next == nil || len(next) == 0 {
			break
		}
		//update levelList
		levelList = next 
	}
	
	return resultList
}

func getLevelNodes(level []*TreeNode) (res []int, next []*TreeNode) {
	for _,l:=range level {
		
		res = append(res, l.Val)

		if l.Left != nil {
			next = append(next, l.Left)
		}
		if l.Right != nil {
			next = append(next, l.Right)
		}
		
	}

	return
}