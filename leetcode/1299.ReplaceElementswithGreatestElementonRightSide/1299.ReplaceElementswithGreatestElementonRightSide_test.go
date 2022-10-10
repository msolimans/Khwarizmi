package replaceelementswithgreatestelementonrightside

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int
        Exp []int
    }{
        {In: []int{17,18,5,4,6,1}, Exp: []int{18,6,6,6,1,-1}},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, replaceElements(test.In))
    }
}