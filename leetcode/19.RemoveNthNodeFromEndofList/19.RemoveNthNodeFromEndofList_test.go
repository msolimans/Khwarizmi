package removenthnodefromendoflist

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)

func TestRemoveNthNodeFromEnd(t *testing.T){
	
	tests := []struct{
		Head *ListNode
		N int  
		Exp *ListNode
	}{
		{Head: Init([]int{1,2,3,4,5}), N: 2, Exp: Init([]int{1,2,3,5})},
		{Head: Init([]int{1,2}), N: 2, Exp: Init([]int{2})},
		{Head: Init([]int{1,2,3}), N: 3, Exp: Init([]int{2,3})},
		{Head: Init([]int{1,2,3,4,5,6}), N: 2, Exp: Init([]int{1,2,3,4,6})},
	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, removeNthFromEnd(test.Head, test.N))
	}

}

func TestRemoveNthNodeFromEnd2(t *testing.T){
	
	tests := []struct{
		Head *ListNode
		N int  
		Exp *ListNode
	}{
		{Head: Init([]int{1,2,3,4,5}), N: 2, Exp: Init([]int{1,2,3,5})},
		{Head: Init([]int{1,2}), N: 2, Exp: Init([]int{2})},
		{Head: Init([]int{1,2,3}), N: 3, Exp: Init([]int{2,3})},
		{Head: Init([]int{1,2,3,4,5,6}), N: 2, Exp: Init([]int{1,2,3,4,6})},
	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, removeNthFromEnd2(test.Head, test.N))
	}

}