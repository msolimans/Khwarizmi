package happynumber

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestCalcSum(t *testing.T) {
	tests := []struct{
		In int 
		Exp int 
	}{
		{In: 5, Exp: 25},
		{In: 15, Exp: 26},

	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, calcSum(test.In))
	}
}


func TestIsHappy(t *testing.T) {
	tests := []struct{
		In int 
		Exp bool 
	}{
		//Explanation:
		// 12 + 92 = 82
		// 82 + 22 = 68
		// 62 + 82 = 100
		// 12 + 02 + 02 = 1
		{In: 19, Exp: true},
		{In: 2, Exp: false},

	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, isHappy(test.In))
	}
}

func TestIsHappy2(t *testing.T) {
	tests := []struct{
		In int 
		Exp bool 
	}{
		//Explanation:
		// 12 + 92 = 82
		// 82 + 22 = 68
		// 62 + 82 = 100
		// 12 + 02 + 02 = 1
		{In: 19, Exp: true},
		{In: 2, Exp: false},

	}

	for _,test := range tests {
		assert.Equal(t, test.Exp, isHappy2(test.In))
	}
}

