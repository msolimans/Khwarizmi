package common

import (
	"container/list"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val int 
	Left *TreeNode 
	Right *TreeNode 
}

//used for unit testing purposes 
func InitTreeNode(input []*int) *TreeNode {
	// [0,2,4,1,null,3,-1,5,1,null,6,null,8]
	root := &TreeNode {Val: *input[0] }
	
	queue := list.New() 
	queue.PushBack(root)

	//[0,2,4,1,null,3,-1,5,1,null,6,null,8]
// 	 0
//    2   4
//  1   3   	
	i := 1
	//not empty 
	for queue.Len() > 0 {
	
		fr := queue.Front()
		queue.Remove(fr)
		node := fr.Value.(*TreeNode)
		if node == nil {
			i++
			continue
		}
		//before we insert make sure we still aligned with array - if not break 
		if i > len(input) - 1 {
			break
		}
		//left and right 
		if input[i] != nil {
			nv := fr.Value.(*TreeNode)
			nv.Left = &TreeNode{Val: *input[i]}
			queue.PushBack(nv.Left)
		}
		i++
		//before we insert make sure we still aligned with array - if not break 
		if i > len(input) - 1 {
			break
		}
		if input[i] != nil {
			nv := fr.Value.(*TreeNode)
			nv.Right = &TreeNode{Val: *input[i]}
			queue.PushBack(nv.Right)
		}
		i++
	}
	
	return root 
 
}

func (t *TreeNode) String() string {
	//use queue to print in BFS 
	queue := list.New()

	queue.PushBack(t)
	var sb strings.Builder
	for queue.Len() > 0 {
		//deq
		f := queue.Front()
		
		if f.Value == nil {
			sb.WriteString("nil-")
			continue
		} 

		node := f.Value.(*TreeNode)
		queue.Remove(f)
		
		if node == nil {
			sb.WriteString("nil-")
			continue
		} 

		//push back left and right 
		sb.WriteString(strconv.Itoa(node.Val) + "-")

		queue.PushBack(node.Left)
		
		queue.PushBack(node.Right)
		
	}

	return sb.String()
}


