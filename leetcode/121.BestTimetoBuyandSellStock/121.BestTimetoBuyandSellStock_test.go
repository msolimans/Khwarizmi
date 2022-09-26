package besttimetobuyandsellstock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	
	tests := []struct{
		In []int
		Exp int 
	} {
		{In: []int{1,2,3,4}, Exp: 3},
		{In: []int{7,1,5,3,6,4}, Exp: 5},
		{In: []int{7,6,5,4,3,2,1}, Exp: 0},
	}
	
	for _, test:= range tests {

		assert.Equal(t, test.Exp, maxProfit(test.In))
	}
}