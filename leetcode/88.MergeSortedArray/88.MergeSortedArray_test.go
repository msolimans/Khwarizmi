package mergesortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	cases := []struct{
		in1 []int
		m int
		in2 []int
		n int 
		out []int 
	}{
		{
			in1: []int{3,4,5, 0, 0 , 0},
			m: 3,
			in2: []int{4,6,7},
			n: 3,
			out: []int{3,4,4,5,6,7},
		},
		{
			in1: []int{5,10,15, 0, 0 , 0},
			m: 3,
			in2: []int{4,6,7},
			n: 3,
			out: []int{4,5,6,7,10,15},
		},
	}
	
	for _, v:=range cases {
		merge(v.in1, v.m, v.in2, v.n)
		assert.Equal(t, v.out, v.in1)
	}
}