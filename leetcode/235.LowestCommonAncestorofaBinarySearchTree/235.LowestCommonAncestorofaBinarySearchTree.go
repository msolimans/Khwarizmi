package lowestcommonancestorofabinarysearchtree

import (
	"container/list"

	. "github.com/msolimans/khawarizmi/leetcode/common"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

 func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pqueue := list.New()
	qqueue := list.New()


	findNode(root, p, pqueue)
	findNode(root, q, qqueue)
	res := root 
	for pqueue.Len() > 0 && qqueue.Len() > 0 {
		pf := pqueue.Front()
		qf := qqueue.Front()

		pfv := pf.Value.(*TreeNode)
		qfv := qf.Value.(*TreeNode)

		pqueue.Remove(pf)
		qqueue.Remove(qf)
		
		if pfv == qfv {
			res = pfv 
		} else {
			break
		}
	}

	return res 
}


func findNode(root, target *TreeNode, queue *list.List) {
	//everytime we traverse to node, we save its ancestor 
	if root == nil {
		return 
	}

	if root == target {
		queue.PushBack(root)
		return 
	}


	//add current node in the path 
	queue.PushBack(root)

	if target.Val <= root.Val {
		findNode(root.Left, target, queue)
	} else {
		findNode(root.Right, target, queue)
	}
}