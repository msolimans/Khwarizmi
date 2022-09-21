package climbingstairs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	assert.Equal(t, 3, climbStairs(3))
	assert.Equal(t,2, climbStairs(2))
}