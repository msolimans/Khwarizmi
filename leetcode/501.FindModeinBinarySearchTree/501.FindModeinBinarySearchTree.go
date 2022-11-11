package findmodeinbinarysearchtree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	"container/list"

	. "github.com/msolimans/khawarizmi/leetcode/common"
)

 func findMode(root *TreeNode) []int {
	 //we can use property of bst for each node find how many do we have and count them accordingly (no extra space)
	 max := 0 
 
	 queue := list.New()
 
	 queue.PushBack(root) 
	 counts := map[int]int{} 
	 res := []int{} 
 
	 for queue.Len() > 0{
		 f := queue.Front()
		 current := f.Value.(*TreeNode)
		 queue.Remove(f)
		 
		 //before count, we use map to store occurences of current node's value (if exists no need to count again)
		 if _, exists := counts[current.Val]; exists {
			 continue //we removed it from queue already 
		 }
 
		 //count current nodes value
		 cnt := countIt(current, current.Val)
		 counts[current.Val] = cnt 
 
		 //we check if current max is less than what we calculated we update max 
		 if max < cnt {
			 //reset result array in this case since the new max (with current node's value) taking place
			 max = cnt 
			 res = []int{} //reset 
			 res = append(res, current.Val)
		 } else if max == cnt {//in case current count is same as current max (add to result array current node)
			 //add to results 
			 res = append(res, current.Val)
		 } //else max is greater than count of current node (no need to do anything here in such case)
 
		 //move to left (in level order traversal)
		 if current.Left != nil {
			 queue.PushBack(current.Left)
		 }
 
		 //move to right 
		 if current.Right != nil {
			 queue.PushBack(current.Right)
		 }
 
	   
	 }
 
	 return res
 
 }
 
 func countIt(root *TreeNode, val int) int {
	 if root == nil {
		 return 0 
	 }
 
	 count := 0 
	 if val == root.Val {
		 count++
	 }
 
	 if val <= root.Val {
		 count += countIt(root.Left, val)
	 } 
	 if val >= root.Val {
		 count += countIt(root.Right, val)
	 }
 
	 return count 
 }
  
