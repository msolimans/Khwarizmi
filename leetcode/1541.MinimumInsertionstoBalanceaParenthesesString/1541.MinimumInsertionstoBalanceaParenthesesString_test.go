package minimuminsertionstobalanceaparenthesesstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinInsertions(t *testing.T) {
    tests := []struct{
        In string 
        Exp int 
    }{
        {In: "(()))", Exp: 1},
        {In: ")", Exp: 2},
        {In: "(", Exp: 2},
        {In: "()", Exp: 1},
        {In: "", Exp: 0},
        {In: "())", Exp: 0},
        {In: "))())(", Exp: 3},
        {In: "))()(()))", Exp: 3},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, minInsertions(test.In))
    }
}