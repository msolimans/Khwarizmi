package twosumivinputisabst

import (
	"container/list"

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



////////////////////////////////////////////////////////////////////////////////////////////////////

//time complexity is still O(logn) 


func findTarget2(root *TreeNode, k int) bool {
	//2 stacks to inject into it ..left and right sides of the tree 
	// all the way down to left represents the lowest elements in the tree
	// all the way down to right represents the biggest elements in the tree 
	//we push both branches and then compare if we get k, return true - if sum is greater than k, we pop from bigger numbers otherwise pop from smaller numbers 

	smallStack := list.New()
	bigStack := list.New()

	//both are equal (stop condition)
	smallStack.PushFront(root)
	bigStack.PushFront(root)

	//build biggest tree 
	buildBig(root, bigStack)
	
	//build smallest tree 
	buildSmall(root, smallStack)
	
	//peek 
	sf := smallStack.Front()
	bf := bigStack.Front()

	for smallStack.Len() > 0 && bigStack.Len() > 0 && sf.Value != bf.Value {
		small := sf.Value.(*TreeNode)
		big := bf.Value.(*TreeNode)

		if small.Val + big.Val == k {
			return true 
		} else if small.Val + big.Val > k {
			//pop from big stack and check the left branch (if we have nodes on it, push them to stack and continue )
			bigStack.Remove(bf)
			if big.Left != nil { //we have more nodes in the left side 
				buildBig(big.Left, bigStack)
			}

			//peek here again 
			bf = bigStack.Front()
		} else {
			//pop from small stack and check the right branch (if we have nodes on it, push them to stack and continue )
			smallStack.Remove(sf)
			if small.Right != nil { //we have more nodes in the right side of this node. before removing, we need those nodes in the stack
				buildSmall(small.Right, smallStack)
			}
			
			//peek here again 
			sf = smallStack.Front()
		}
	}

	return false 


}

func buildBig(root *TreeNode, stack *list.List)  {
	c := root
	for c != nil && c.Left != nil  {
		stack.PushFront(c.Left) //we always insert non-nil values 
		c = c.Left
	}
}

func buildSmall(root *TreeNode, stack *list.List)  {
	
	c := root 
	
	for c != nil && c.Right != nil {
		stack.PushFront(c.Right) //we always insert non-nil values 
		c = c.Right
	}
}


//pop from biggest tree 
//pop from smallest tree 