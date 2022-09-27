package reversewordsinastringiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T){
	tests := []struct{
		In string 
		In2 int 
		Exp string 
	}{
		{In: "abcdefg", In2: 2, Exp: "bacdfeg"},
		{In: "abcdefg", In2: 2, Exp: "bacdfeg"},
		{In: "abcdefgk", In2: 2, Exp: "bacdfegk"},
		{In: "abcdefgk", In2: 3, Exp: "cbadefkg"},
		{In: "abcd", In2: 2, Exp: "bacd"},
	}

	for _,test:= range tests {
		assert.Equal(t, test.Exp, reverseStr(test.In, test.In2))
	}
}