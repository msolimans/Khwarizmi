package checkifstringisaprefixofarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In string 
		Words []string 
        Exp bool 
    }{
        {In: "iloveleetcode", Words: []string{"i","love","leetcode","apples"}, Exp: true },
        {In: "iloveleetcode", Words: []string{"apples","i","love","leetcode"}, Exp: false },
        {In: "cccccccccccc", Words: []string{"c","cc"}, Exp: false },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,isPrefixString(test.In, test.Words))
    }
}
