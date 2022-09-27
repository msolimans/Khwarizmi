package singlenumberiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct{
		In []int
		Exp int 
	} {
		{In: []int{1,3,3,1,1,2,3}, Exp: 2},
		{In: []int{0,0,1,0}, Exp: 1},
		{In: []int{0,0,0, 1}, Exp: 1},
		{In: []int{0,1,1, 1}, Exp: 0},
	}

	for _, test := range tests {
		res := singleNumber(test.In)
		assert.Equal(t, test.Exp, res)
	}
}

