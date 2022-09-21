package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T){
	tests := []struct{
		In1 []int 
		In2 int 
		Exp int 
	}{
		{[]int{-1,0,3,5,9,12}, 9, 4},
		{[]int{-1,0,3,5,9,12}, 2, -1},
		{[]int{1}, -1, -1},
		{[]int{1,2}, 3, -1},
		{[]int{1,2,3,10}, 10, 3},
		{[]int{1,2,3,10}, 1, 0},
	}

	for _, test := range tests {
		assert.Equal(t, test.Exp, search(test.In1, test.In2))
	}
}