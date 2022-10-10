package checkifallasappearsbeforeallbs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In string 
        Exp bool 
    }{
        {In: "aaabbb", Exp: true },
        {In: "abab", Exp: false },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,checkString(test.In))
    }
}