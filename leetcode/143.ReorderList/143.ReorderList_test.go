package reorderlist

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestReorderList(t *testing.T) {
	tests := []struct{
		In []int
		Exp []int 
	} {
		{In: []int{1,2,3,4}, Exp: []int{1,4,2,3}},
		{In: []int{1,2,3,4,5}, Exp: []int{1,5,2,4,3}},
		{In: []int{1,2}, Exp: []int{1,2}},
		{In: []int{1,2,3}, Exp: []int{1,3,2}},
	}
	
	for _, test:= range tests {
		list := Init(test.In)
		reorderList(list )

		assert.Equal(t, Init(test.Exp), list)
	}
}

func TestInsertAfter(t *testing.T){
	tests := []struct{
		In []int
		Skip int  
		Exp []int 
	} {
		{In: []int{1,2,3,4,5}, Skip: 1, Exp: []int{1,4,3,2,5}},
		{In: []int{1,2,3,4,5},Skip: 2, Exp: []int{2,1,3,4,5}},
	}

	for _, test := range tests {

		what := &ListNode{Val: 1000}
		list := Init(test.In)
		c := list 
		for i := 0 ;i<test.Skip;i++ {
			if i == test.Skip {
				insertAfter(list, what)
				assert.Equal(t,1000, c.Next.Val)
			}
			c = c.Next
		}
	}
}


func TestFindMedium(t *testing.T) {
	tests := []struct{
		In []int
		Exp1 int 
		Exp2 int 
	} {
		{In: []int{1,2,3,4,5},  Exp1: 2, Exp2: 3},
		{In: []int{1,2,3,4},  Exp1: 2, Exp2: 3},
		{In: []int{1,2,3,4,5, 6},  Exp1: 3, Exp2: 4},
		{In: []int{1},  Exp1: 1, Exp2: 1},
		{In: []int{1,2},  Exp1: 1, Exp2: 2},
	}

	for _, test := range tests {
		pmed, med := findMedium(Init(test.In)) 
		assert.Equal(t, test.Exp1, pmed.Val)
		assert.Equal(t, test.Exp2, med.Val)
	}
}


