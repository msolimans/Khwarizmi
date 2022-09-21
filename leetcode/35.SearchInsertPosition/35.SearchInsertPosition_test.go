package searchinsertposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	tests  := []struct {
		In1 []int
		In2 int 
		Exp int 
	} {
		{[]int{1,3}, 2, 1},
		{[]int{1,3}, 2, 1},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, 1},
		{[]int{2}, 1, 0},
		{[]int{1,3,5,6}, 5, 2},
		{[]int{1,3,5,6}, 2, 1},
		{[]int{1,3,5,6}, 7, 4},
	}

	for _, test := range tests {
		assert.Equal(t, test.Exp, searchInsert(test.In1, test.In2))
	}
	 
}