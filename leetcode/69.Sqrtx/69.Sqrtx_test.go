package sqrtx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrtx(t *testing.T){
	assert.Equal(t, 2, mySqrt(4))
	assert.Equal(t, 2, mySqrt(8))
	assert.Equal(t, 3, mySqrt(9))
}