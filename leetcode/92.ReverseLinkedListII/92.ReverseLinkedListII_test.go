package reverselinkedlistii

import (
	"testing"

	. "github.com/msolimans/khawarizmi/leetcode/common"
	"github.com/stretchr/testify/assert"
)


func TestReverseBetween(t *testing.T){
	tests := []struct{
		In []int 
		Left int
		Right int
		Exp []int 
	} {
		{In: []int{1,2,3,4,5}, Left: 2, Right: 4, Exp: []int{1,4,3,2,5}},
		{In: []int{1,2,3,4,5}, Left: 1, Right: 2, Exp: []int{2,1,3,4,5}},
		{In: []int{1,2,3,4,5}, Left: 3, Right: 5, Exp: []int{1,2,5,4,3}},
		{In: []int{5}, Left: 1, Right: 1, Exp: []int{5}},
	}

	for _, test := range tests {
		res := reverseBetween(Init(test.In), test.Left, test.Right)
		assert.Equal(t, Init(test.Exp), res)
	}
	
}



func TestReverse(t *testing.T){
	tests := []struct{
		In []int 
		Exp1 int 
		Exp2 int 
	} {
		{In: []int{1,2,3,4}, Exp1: 4, Exp2: 1},
		{In: []int{1,2,3}, Exp1: 3, Exp2: 1},
		{In: []int{1,2}, Exp1: 2, Exp2: 1},
		{In: []int{1}, Exp1: 1, Exp2: 1},
	}

	for _, test := range tests {
		rhead, rtail := reverse(Init(test.In))
	
		assert.Equal(t, test.Exp1, rhead.Val)
		assert.Equal(t, test.Exp2, rtail.Val)
	}
	
}


