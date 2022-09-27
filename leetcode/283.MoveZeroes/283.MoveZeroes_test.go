package movezeroes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeros(t *testing.T) {
	tests := []struct{
		In []int 
		Exp []int 
	}{
		{[]int{1,2,0,0,3,4}, []int{1,2,3,4,0,0}},
		{[]int{0,0,1,2,3,4}, []int{1,2,3,4,0,0}},
		{[]int{2,3,4,0,0,0}, []int{2,3,4,0,0,0}},
		{[]int{2,3,4}, []int{2,3,4}},
	}

	for _,test:=range tests {
		moveZeroes(test.In)
		assert.Equal(t, test.Exp, test.In)
	}
}



func TestMoveZeros2(t *testing.T) {
	tests := []struct{
		In []int 
		Exp []int 
	}{
		{[]int{1,2,0,0,3,4}, []int{1,2,3,4,0,0}},
		{[]int{0,0,1,2,3,4}, []int{1,2,3,4,0,0}},
		{[]int{2,3,4,0,0,0}, []int{2,3,4,0,0,0}},
		{[]int{2,3,4}, []int{2,3,4}},
	}

	for _,test:=range tests {
		moveZeroes2(test.In)
		assert.Equal(t, test.Exp, test.In)
	}
}