package onlinestockspan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        In []int 
        Exp []int 
    }{
        {In: []int{100, 80, 60, 70, 60, 75, 85} , Exp: []int{ 1, 1, 1, 2, 1, 4, 6}},
    }

    for _, test := range tests {
		stock := Constructor()
		for i, v := range test.In {
			assert.Equal(t, test.Exp[i],	stock.Next(v))
		}
       
    }
}
