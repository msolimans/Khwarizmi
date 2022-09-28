package removenthnodefromendoflist

import . "github.com/msolimans/khawarizmi/leetcode/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func removeNthFromEnd(head *ListNode, n int) *ListNode {
    //easy: momve count-n position then delete that node 
    //how many elems in the linkedList 
    
    c := head
    count := 0
    
    for c != nil {
        count++
        c = c.Next 
    }
    
    if count == 1 {
        // head = nil 
        return nil
    }
    
    c = head 
    i := 1
    
    for i < count-n {
        c = c.Next 
        i++
    }
    
    // 1 < 3    2 < 3 
    // 2        3
    // 2        3 
    
    //remove next 

	if i == 1 && n == count { //first item to remove (don't remove next in this case) handling cases like [1,2] and n=2  
		return c.Next
	}
    
    c.Next = c.Next.Next 
    return head 
}


//2nd solutions (first uses 2 passes while this one uses 1 pass)
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	preHead := &ListNode{Next: head}

	//use slow and fast pointers -> 
	//suppose size is number of elements in linked list while n is the index where we need to remove
	//we will use 2 pointers 
	//slow pointer where it stays at head without being moved 
	//fast pointer where we start moving to index "n" now the left over of elements in the list is "size-n"
	//when the fast pointer points to "n", we start moving "slow" until fast becomes nil 
	//now slow points at "size-n"

	fast, slow := preHead, preHead

	for i :=0;i<n;i++ {
		fast = fast.Next
	}

	//now fast points at n, start moving 
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	//now remove the next from slow 
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}
	
	return preHead.Next
}