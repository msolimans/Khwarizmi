package reversevowelsofastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In string 
        Exp string 
    }{
        {In: "hello", Exp: "holle"},
        {In: "leetcode", Exp: "leotcede"},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, reverseVowels(test.In))
    }
}