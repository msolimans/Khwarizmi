package MergeTwoSortedLists

import . "github.com/msolimans/khawarizmi/leetcode/common"

func copyFrom(list *ListNode, c *ListNode) (*ListNode, *ListNode) {
	c.Next = &ListNode{Val: list.Val}
	c = c.Next
	list = list.Next
	return list, c
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	preHead := &ListNode{}
	c := preHead

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			list1, c = copyFrom(list1, c)
		} else {
			list2, c = copyFrom(list2, c)
		}
	}

	for list1 != nil {
		list1, c  = copyFrom(list1, c)
	}

	for list2 != nil {
		list2, c  = copyFrom(list2, c)
	}

	return preHead.Next
}

