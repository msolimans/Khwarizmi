package singlenumberiii

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct{
		In []int
		Exp []int 
	} {
		{In: []int{1,3,1,2}, Exp: []int{2,3}},
		{In: []int{0,-1}, Exp: []int{-1,0}},
	}

	for _, test := range tests {
		res := singleNumber(test.In)
		sort.Ints(res)
		assert.Equal(t, test.Exp, res)
	}
}


func TestSingleNumber2(t *testing.T) {
	tests := []struct{
		In []int
		Exp []int 
	} {
		{In: []int{1,3,1,2}, Exp: []int{2,3}},
		{In: []int{0,-1}, Exp: []int{-1,0}},
		{In: []int{5,0,-1}, Exp: []int{-1,0,5}},
	}

	for _, test := range tests {
		res := singleNumber2(test.In)
		sort.Ints(res)
		assert.Equal(t, test.Exp, res)
	}
}