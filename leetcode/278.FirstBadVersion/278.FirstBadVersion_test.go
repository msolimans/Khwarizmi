package firstbadversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		In int 
		Ex int 
	}{
		////bad versions are 12 
		{16, 12},
	}

	for _, test := range tests {
		assert.Equal(t, test.Ex, firstBadVersion(test.In))

	}
}



func TestIsBadVersion(t *testing.T){
	assert.Equal(t, false, isBadVersion(1))
	assert.Equal(t, true, isBadVersion(12))
	
}