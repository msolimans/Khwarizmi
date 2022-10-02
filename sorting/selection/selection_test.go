package selection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	tests := []struct {
		In []int 
		Exp []int 
	}{
		{In: []int{1,2,1,9,2,8}, Exp: []int{1,1,2,2,8,9}},
		{In: []int{8,6,4,2,1}, Exp: []int{1,2,4,6,8}},
		{In: []int{8}, Exp: []int{8}},
		{In: []int{9,1}, Exp: []int{1,9}},
	}

	for _,test:= range tests {
		assert.Equal(t, test.Exp, sort(test.In))
	}
}