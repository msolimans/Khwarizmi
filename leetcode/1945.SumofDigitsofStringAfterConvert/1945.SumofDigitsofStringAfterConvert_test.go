package sumofdigitsofstringafterconvert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLucky(t *testing.T){
	tests := []struct{
        In string 
		K int 
        Exp int 
    }{
        {In: "hvmhoasabaymnmsd", K: 1, Exp: 79},
		//h v > mhoasabaymnmsd
		//6+0+4+5+5+5+1+2+2+9+2+5+9+5+2+8
        {In: "a", K: 1, Exp: 1},
        {In: "iiii", K: 1, Exp: 36},
        {In: "leetcode", K: 2, Exp: 6},
        {In: "zbax", K: 2, Exp: 8},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, getLucky(test.In, test.K))
    }
}