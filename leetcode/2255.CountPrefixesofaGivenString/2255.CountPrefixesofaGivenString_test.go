package countprefixesofagivenstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []string
		S string 
        Exp int 
    }{
        {In: []string{"a","b","c","ab","bc","abc"}, S: "abc", Exp: 3 },
        {In: []string{"a","a"}, S: "aa", Exp: 2 },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,countPrefixes(test.In, test.S))
    }
}
