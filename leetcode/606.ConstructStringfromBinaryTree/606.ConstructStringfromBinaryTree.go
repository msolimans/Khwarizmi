package constructstringfrombinarytree

import (
	"strconv"
	"strings"

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
 func tree2str(root *TreeNode) string {
    sb = strings.Builder{}
    preorder(root)
    res := sb.String()
    return res[1:len(res) - 1]
}   

var sb strings.Builder 

func preorder(root *TreeNode) {
    if root == nil {
        sb.WriteString("()")
        return
    }
    
    sb.WriteString("(")
    sb.WriteString(strconv.Itoa(root.Val))
    
    if (root.Left != nil && root.Right != nil) || (root.Left == nil && root.Right != nil)  {
        preorder(root.Left)
        preorder(root.Right)
    } else if root.Left != nil && root.Right == nil {
          preorder(root.Left)
    }
    
    sb.WriteString(")")
    
}