package sumofroottoleafbinarynumbers

import (
	"strconv"

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
 func sumRootToLeaf(root *TreeNode) int {
    res = []string{}
    combineNumbers(root, "")
	//loop thru combined numbers and sum them 
    sum := 0 
    for _, v := range res {
		//parse as int
        i, _ := strconv.ParseInt(v, 2, 32)
        sum += int(i) 
    }

    return sum 
}

var res []string 

func combineNumbers(root *TreeNode, s string)  {
	//in case (left here is nil so if we dun have below condition it will fail)
	// 	1
	//	  \
	//		1

	if root == nil {
		return 
	}

	//if this a leaf node, will add prev + current value  
	if root.Left == nil && root.Right == nil {
		res = append(res, s+strconv.Itoa(root.Val))
		return
	}

    combineNumbers(root.Left, s+strconv.Itoa(root.Val))
    combineNumbers(root.Right, s+strconv.Itoa(root.Val))   
}