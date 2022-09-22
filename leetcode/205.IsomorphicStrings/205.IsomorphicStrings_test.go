package isomorphicstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
var tests = []struct{
	In1 string 
	In2 string 
	Exp bool 
} {
	{"egg","add",true },
	{"foo","bar",false },
	{"paper","title",true },
}

func TestIsIsomorphic(t *testing.T) {


	for _,test := range tests {
		assert.Equal(t, test.Exp, isIsomorphic(test.In1, test.In2))
		
	}

}
func TestIsIsomorphic2(t *testing.T) {
	for _,test := range tests {
		assert.Equal(t, test.Exp, isIsomorphic(test.In1, test.In2))
		
	}
}

func TestTransform(t *testing.T) {
	tests := []struct{
		In string 
		Exp string 
	} {
		{"addat", "01104"},
		{"errea", "01104"},
		{"errsa", "01134"},
		{"paper", "01034"},
		{"titee", "01033"},
	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, transform(test.In))
	}

}