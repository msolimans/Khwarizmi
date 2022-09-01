package MergeTwoSortedLists

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T){
	list1 := &ListNode{Val: 10}
	list1.Next = &ListNode{Val: 15}
	list1.Next.Next = &ListNode{Val: 20}
	list2 := &ListNode{Val: 8}
	list2.Next = &ListNode{Val: 9}
	list2.Next.Next = &ListNode{Val: 14}
	
	out := &ListNode{Val: 8}
	out.Next = &ListNode{Val: 9}
	out.Next.Next = &ListNode{Val: 10}
	out.Next.Next.Next = &ListNode{Val: 14}
	out.Next.Next.Next.Next = &ListNode{Val: 15}
	out.Next.Next.Next.Next.Next = &ListNode{Val: 20}
 
	// []string{}
	cases := []struct{
		in1 *ListNode 
		in2 *ListNode 
		out *ListNode 
	}{
		{
			in1: list1,
			in2: list2,
			out: out ,
		},
		{ 
			in1: nil,
			in2: nil,
			out: nil,
		},
		{
			in1: &ListNode{Val: 10},
			in2: nil,
			out: &ListNode{Val: 10},
		},
		{
			in1: nil,
			in2: &ListNode{Val: 10},
			out: &ListNode{Val: 10},
		},
	}

	for _,v :=range cases {
		assert.True(t, mergeTwoLists(v.in1, v.in2).Equals(v.out))
	}
}