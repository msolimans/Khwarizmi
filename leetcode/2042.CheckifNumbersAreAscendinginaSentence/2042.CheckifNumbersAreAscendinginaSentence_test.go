package checkifnumbersareascendinginasentence

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In string 
        Exp bool 
    }{
        {In: "1 box has 3 blue 4 red 6 green and 12 yellow marbles", Exp: true},
        {In: "hello world 5 x 5", Exp: false},
        {In: "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s", Exp: false},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,areNumbersAscending(test.In))
    }
}