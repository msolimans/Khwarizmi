package removealladjacentduplicatesinstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T){
	tests := []struct{
		In string
		Exp string 
	} {
		{In: "abbaca", Exp: "ca"},
		{In: "azxxzy", Exp: "ay"},
		{In: "azxzxzy", Exp: "azxzxzy"},
	}

	for _,test:= range tests {
		assert.Equal(t, test.Exp, removeDuplicates(test.In))
	}

}