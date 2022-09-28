package backspacestringcompare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackSpaceCompare(t *testing.T){
	tests := []struct{
		In string 
		In2 string 
		Exp bool 
	} {
		{In: "ab#c", In2: "ad#c", Exp: true },
		{In: "ab##", In2: "c#d#", Exp: true },
		{In: "a#c", In2: "b", Exp: false },
		{In: "a##c", In2: "#a#c", Exp: true },
		{In: "bxj##tw", In2: "bxo#j##tw", Exp: true },
		{In: "a#b#c#da##cs#", In2: "e#f#ew##r#bs##", Exp: false },
		{In: "y#fo##f", In2: "y#f#o##f", Exp: true },
	 
	}

	for _,test:= range tests {
		assert.Equal(t, test.Exp, backspaceCompare(test.In,test.In2))
	}
}