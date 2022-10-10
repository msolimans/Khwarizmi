package breakapalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In string 
        Exp string 
    }{
        {In: "aba", Exp: "abb"},
        {In: "bab", Exp: "aab"},
        {In: "a", Exp: ""},
        {In: "abccba", Exp: "aaccba"},
        {In: "bbbbaaaa", Exp: "abbbaaaa"},
        {In: "aaaa", Exp: "aaab"},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, breakPalindrome(test.In))
    }
}