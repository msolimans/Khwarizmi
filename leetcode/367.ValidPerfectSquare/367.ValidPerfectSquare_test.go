package validperfectsquare

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPerfectSquare(t *testing.T) {
	assert.True(t, isPerfectSquare(16))
	assert.False(t, isPerfectSquare(14))
	assert.True(t, isPerfectSquare(1))
}