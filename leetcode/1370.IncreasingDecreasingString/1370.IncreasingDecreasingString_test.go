package increasingdecreasingstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In string 
        Exp string 
    }{
        {In: "aaaabbbbcccc", Exp: "abccbaabccba"},
        {In: "rat", Exp: "art"},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,sortString(test.In))
    }
}

