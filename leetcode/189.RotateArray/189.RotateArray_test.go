package rotatearray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T){
	tests := []struct{
		In1 []int 
		In2 int
		Exp []int
	} {
		{[]int{1,2,3,4,5,6,7}, 3, []int{5,6,7,1,2,3,4}},
		{[]int{-1,-100,3,99}, 2,  []int{3,99,-1,-100}},
		{[]int{-1,2}, 6,  []int{-1,2}},
		{[]int{-1}, 2,  []int{-1}},
	}

	for _,test := range tests {
		rotate(test.In1, test.In2)
		assert.Equal(t, test.Exp, test.In1)
	}
}