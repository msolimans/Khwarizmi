package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func  TestIterativeString(t *testing.T) {
	tests := []struct{
		In *Node  
		Exp string 
	}{
		{
			In: &Node{
				Val: 1,
				Children:[]*Node{{Val: 2}, {Val: 3}},
			},
			Exp: "123",
		},
	}
	for _,test := range tests{
		assert.Equal(t, test.Exp, test.In.ToString())
	}
}