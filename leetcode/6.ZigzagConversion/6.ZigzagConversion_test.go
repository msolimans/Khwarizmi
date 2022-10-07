package zigzagconversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
    tests := []struct{
        In string 
		NumRows int 
        Exp string 
    }{
        {In: "PAYPALISHIRING", NumRows: 3, Exp: "PAHNAPLSIIGYIR"},
        {In: "PAYPALISHIRING", NumRows: 4, Exp: "PINALSIGYAHRPI"},
        {In: "A", NumRows: 1, Exp: "A"},
        {In: "AB", NumRows: 1, Exp: "AB"},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, convert(test.In, test.NumRows))
    }
}


