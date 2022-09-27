package reversewordsinastringiii

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T){
	tests := []struct{
		In string 
		Exp string 
	}{
		{In: "Let's take LeetCode contest", Exp: "s'teL ekat edoCteeL tsetnoc"},
		{In: "God Ding", Exp: "doG gniD"},
	}

	for _,test:= range tests {
		assert.Equal(t, test.Exp, reverseWords(test.In))
	}
}