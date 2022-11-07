package maximum69number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In int 
        Exp int 
    }{
        {In: 669, Exp: 969 },
        {In: 9996, Exp: 9999 },
        {In: 9999, Exp: 9999 },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,maximum69Number(test.In))
    }
}
