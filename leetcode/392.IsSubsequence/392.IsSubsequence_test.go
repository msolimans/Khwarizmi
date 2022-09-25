package issubsequence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubsequence(t *testing.T) {
	var tests = []struct{
		In1 string 
		In2 string 
		Exp bool 
	} {
		{"abc","ahbgdc",true },
		{"axc","ahbgdc",false },
		{"ahbgdc", "abc",false },
		// {"paper","title",true },
	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, isSubsequence(test.In1, test.In2))
	}

}