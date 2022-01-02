package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWrapStringArray(t *testing.T) {
	type test struct {
		array    []string
		expected []interface{}
	}
	tests := []test{
		{
			array:    []string{},
			expected: []interface{}{},
		},
		{
			array:    []string{"hello", "world"},
			expected: []interface{}{"hello", "world"},
		},
	}
	for _, tst := range tests {
		for i, a := range tst.array {
			e := tst.expected[i]
			assert.Equal(t, e.(string), a)
		}
	}
}
