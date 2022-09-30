package calculatedigitsumofastring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigitSum(t *testing.T) {
    tests := []struct{
        In string 
		K int 
        Exp string 
    }{
        {In: "11111222223", K: 3, Exp: "135"},
        {In: "00000000", K: 3, Exp: "000"},
		
		      
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, digitSum(test.In, test.K))
    }
}


func TestSum(t *testing.T) {
    tests := []struct{
        In string 
		Exp string 
    }{
        {In: "12", Exp: "3"},
        {In: "1234", Exp: "10"},
		
		      
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp, sum(test.In))
    }
}
