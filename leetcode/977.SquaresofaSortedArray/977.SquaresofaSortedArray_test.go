package squaresofasortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct{
		In []int 
		Exp []int 
	} {
		{[]int{-4,-1,0,3,10}, []int{0,1,9,16,100}},
		{[]int{-7,-3,2,3,11}, []int{4,9,9,49,121}},
	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, sortedSquares(test.In))
	}
}