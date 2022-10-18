package countandsay

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In int 
        Exp string 
    }{
        {In: 1, Exp: "1"},
        {In: 4, Exp: "1211"},
        {In: 10, Exp: "13211311123113112211"},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, countAndSay(test.In))
    }
}
