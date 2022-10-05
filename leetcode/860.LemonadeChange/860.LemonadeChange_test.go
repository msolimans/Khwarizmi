package lemonadechange

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLemonadeChange(t *testing.T) {
    tests := []struct{
        In []int  
        Exp bool
    }{
        {In: []int{5,5,5,10,20}, Exp: true  },
        {In: []int{5,5,10,10,20}, Exp: false  },
        {In: []int{10,5}, Exp: false  },
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, lemonadeChange(test.In))
    }
}
