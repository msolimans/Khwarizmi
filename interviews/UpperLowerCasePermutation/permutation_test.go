package permutation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
    tests := []struct{
        Word string 
        CharSet string 
		Exp []string 
    }{
        {Word: "medium-one", CharSet: "dn" , Exp: []string{"medium-one", "meDium-one", "meDium-oNe", "medium-oNe"}},
        {Word: "medium-one", CharSet: "md" , Exp: []string{"medium-one", "Medium-one", "MeDium-one", "MeDiuM-one", "meDium-one", "meDiuM-one", "mediuM-one"}},
        {Word: "medium-one", CharSet: "m" , Exp: []string{"medium-one", "Medium-one", "MediuM-one", "mediuM-one"}},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,permutateWord(test.Word, test.CharSet))
    }
}



func TestExtractWord(t *testing.T) {
    tests := []struct{
        Word string 
		Indices []int
		S int 
		E int 
		Exp string 
    }{
        {Word: "medium-one", Indices: []int{1,2,3,4}, S: 0, E: 2, Exp: "mEDIum-one"},
    }

    for _, test := range tests {
        assert.Equal(t, test.Exp,extractWord([]byte(test.Word), test.Indices, test.S, test.E))
    }
}

//1 2 3 
//1 
//2 
//3 
//1 2 
//1 3 
//2 3 
//1 2 3 

//1 2 3 4 
//1 2 3 4 sep
//1 2 => 1 3 => 1 4 
//1 3 => 2 4 
//1 

