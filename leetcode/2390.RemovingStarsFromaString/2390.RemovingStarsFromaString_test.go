package removingstarsfromastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveStars(t *testing.T) {
	tests := []struct{
		In string 
		Exp string 
	}{
		{In: "leet**cod*e", Exp: "lecoe"},
		{In: "erase*****", Exp: ""},
	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, removeStars(test.In))
	}
}