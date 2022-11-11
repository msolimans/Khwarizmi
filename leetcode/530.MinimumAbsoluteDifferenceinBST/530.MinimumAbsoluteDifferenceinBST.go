package minimumabsolutedifferenceinbst

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

//similar to isValidBST (boundaries for each node and calc distance between both side and store the min distance)

func getMinimumDifference(root *TreeNode) int {
    res = math.MaxFloat64 
    explore(root, math.MinInt, math.MaxInt)
    return int(res)
}

var res float64 = math.MaxFloat64 

func explore(root *TreeNode, min, max int) {
    if root == nil {
        return
    }

    explore(root.Left, min, root.Val)
    explore(root.Right, root.Val, max)
    
    leftDist := math.Abs(float64(root.Val - min))
    rightDist := math.Abs(float64(root.Val - max))

    if leftDist <= rightDist && leftDist < res {
        res = leftDist 
    } else if rightDist < leftDist && rightDist < res {
        res = rightDist 
    }
}