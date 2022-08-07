package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	set := make(Set[int], 0)

	set.Add(1)
	_, ok := set[1]
	assert.True(t, ok)

	_, ok = set[2]
	assert.False(t, ok)
}

func TestRemove(t *testing.T) {

	set := make(Set[int], 0)

	set.Remove(1)
}
