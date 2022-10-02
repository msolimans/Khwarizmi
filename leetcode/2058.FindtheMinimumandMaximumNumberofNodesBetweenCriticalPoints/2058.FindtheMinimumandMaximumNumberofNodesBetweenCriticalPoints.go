package findtheminimumandmaximumnumberofnodesbetweencriticalpoints

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 func nodesBetweenCriticalPoints(head *ListNode) []int {
    //identify/find critical points 
    //record indexes 
    //biggest - smallest => gives the maxDistance 
    //1 2 3 4 => sort them => biggest - smallest and compare in between indexes 
    
    if head == nil || head.Next == nil || head.Next.Next == nil {
        return []int{-1,-1}
    }
    
    fast := head.Next 
    slow := head 
    i := 2
    indexes := []int{}

	//2 pointers to maintain 2 positions 
	//slow represents the previous node to the current 
	//fast represents the current node 
	//we save into indexes the positions of critical points or nodes 
    for fast != nil && fast.Next != nil {
        
        if fast.Val < slow.Val && fast.Val < fast.Next.Val {
            indexes = append(indexes, i)
        }
        
        if fast.Val > slow.Val && fast.Val > fast.Next.Val {
            indexes = append(indexes, i)
        }
        
        i++
        slow = slow.Next 
        fast = fast.Next 
    }
    
	//if we got 0 or 1 in indexes, we return -1,-1 
    if len(indexes) == 0 {
        return []int{-1,-1}
    } else if len(indexes) == 1 {
        return []int{-1,-1}
    } else if len(indexes) == 2 {
        return []int{indexes[1]-indexes[0], indexes[1]-indexes[0]}
    }
    
    //since indexes are already sorted, we subtract the first index from the last to get the max distance 
    max:= indexes[len(indexes) - 1] - indexes[0]
    
    //find minimal distance (needs comparisons for every index we recorded) 
    min := indexes[1] - indexes[0]
     
    for i:=1;i<len(indexes);i++ {
        if indexes[i] - indexes[i-1] < min {
            min = indexes[i] - indexes[i-1]
        }
    }
    
    return []int{min, max} 
}

 