package narytreepreordertraversal

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */


func preorder(root *Node) []int {
	res := []int{}
	iPreorder(root, &res)
	return res 
}

func iPreorder(root *Node, res *[]int) {
	if root == nil {
		return 
	}

	*res = append(*res, root.Val)

	for _,c := range root.Children {
		iPreorder(c, res)
	}
}


