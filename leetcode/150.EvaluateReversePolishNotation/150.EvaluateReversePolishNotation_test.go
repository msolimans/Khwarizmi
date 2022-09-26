package evaluatereversepolishnotation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvalRPN(t *testing.T){
	tests := []struct{
		In []string 
		Exp int 
	} {
		// {In: []string{"1", "2", "+", "3","*"}, Exp: 9},
		{In: []string{"1", "2", "+", "3","*", "3", "/"}, Exp: 3},
		{In: []string{"1", "2", "+", "3","*", "3", "/", "2","-"}, Exp: 1},
	}
	
	for _, test:= range tests {
		assert.Equal(t, test.Exp, evalRPN(test.In))

	}
}