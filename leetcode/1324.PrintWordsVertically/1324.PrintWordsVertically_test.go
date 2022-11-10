package printwordsvertically

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In string 
        Exp []string 
    }{
        {In: "HOW ARE YOU", Exp: []string{"HAY","ORO","WEU"}},
        {In: "TO BE OR NOT TO BE", Exp: []string{"TBONTB","OEROOE","   T"}},
        {In: "CONTEST IS COMING", Exp: []string{"CIC","OSO","N M","T I","E N","S G","T"}},

    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,printVertically(test.In))
    }
}
