package containsduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
    tests := []struct{
        In []int
        Exp bool
    }{
        {In: []int{1,2,3,1}, Exp: true},
        {In: []int{2,3,1}, Exp: false},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, containsDuplicate(test.In))
    }
}