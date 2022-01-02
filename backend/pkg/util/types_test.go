package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBool(t *testing.T) {
	assert.Equal(t, *Bool(true), true)
	assert.Equal(t, *Bool(false), false)
}

func TestInt(t *testing.T) {
	for _, num := range []int{0, 1, 2, -1, -2} {
		assert.Equal(t, *Int(num), num)
	}
}
