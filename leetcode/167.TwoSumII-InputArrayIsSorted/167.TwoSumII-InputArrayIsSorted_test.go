package twosumiiinputarrayissorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T){
	tests := []struct{
		In []int 
		In2 int 
		Exp []int 
	}{
		{In: []int{2,7,11,15}, In2: 9, Exp: []int{1,2}},
		{In: []int{-1,0}, In2: -1, Exp: []int{1,2}},
		{In: []int{2,3,4}, In2: 6, Exp: []int{1,3}},
	}

	for _, test := range tests {
		assert.Equal(t, test.Exp, twoSum(test.In, test.In2))
	}
}