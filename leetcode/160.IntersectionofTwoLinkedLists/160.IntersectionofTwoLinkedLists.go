package intersectionoftwolinkedlists

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
import (
	"container/lists"

	. "github.com/msolimans/khawarizmi/leetcode/common"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	st1 := make(lists.List)
	st2 := make(lists.List)

	for headA != nil {
		st1.PushFront(headA)
		headA = headA.Next
	}

	for headB != nil {
		st2.PushFront(headB) 
		headB = headB.Next
	}

	var prev *ListNode = nil 
	for st1.Front() != nil && st2.Front() != nil {
		if st1.Front().Value == st2.Front().Value {
			prev = st1.Front().Value
			st1.Remove(st1.Front())
			st2.Remove(st2.Front())
		} else { //not equal (here we can break and return previous intersection point)
			return prev
		}
	}

	return nil
}