package removeduplicatesfromsortedarray

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicaes(t *testing.T) {
	assert.Equal(t,4, removeDuplicates([]int{1,2,2,3,4}))
}