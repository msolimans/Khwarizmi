package rotatestring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateString(t *testing.T) {
    tests := []struct{
        In string 
		In2 string 
        Exp bool
    }{
        {In: "abcde", In2: "cdeab", Exp: true },
        {In: "abcde", In2: "abced", Exp: false },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, rotateString(test.In, test.In2))
    }
}
