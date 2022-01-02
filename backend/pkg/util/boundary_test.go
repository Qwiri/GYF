package util

import "testing"

func i(i ...int) []int {
	return i
}

func TestBoundaries_Applies(t *testing.T) {
	type test struct {
		bounds   Boundaries
		input    []int
		expected bool
	}
	tests := []test{
		// single exact
		{
			bounds:   Bounds(BoundExact(1)),
			input:    i(1),
			expected: true,
		},
		{
			bounds:   Bounds(BoundExact(1)),
			input:    i(2, 0),
			expected: false,
		},
		// single min
		{
			bounds:   Bounds(BoundMin(2)),
			input:    i(3),
			expected: true,
		},
		{
			bounds:   Bounds(BoundMin(2)),
			input:    i(0, -3),
			expected: false,
		},
		// single max
		{
			bounds:   Bounds(BoundMax(2)),
			input:    i(0, 1, 2),
			expected: true,
		},
		{
			bounds:   Bounds(BoundMax(2)),
			input:    i(3),
			expected: false,
		},
		// multi exact
		{
			bounds:   Bounds(BoundExact(1), BoundExact(2)),
			input:    i(0, 1, -1),
			expected: false,
		},
		// between
		{
			bounds:   Bounds(BoundMin(5), BoundMax(10)),
			input:    i(5, 6, 7, 8, 9, 10),
			expected: true,
		},
		{
			bounds:   Bounds(BoundMin(5), BoundMax(10)),
			input:    i(4, 11, -4, -11),
			expected: false,
		},
	}
	for _, tst := range tests {
		for _, in := range tst.input {
			got := tst.bounds.Applies(in)
			if got != tst.expected {
				t.Errorf("Boundaries() bounds: %v, input: %v, expected: %v got %v", tst.bounds, in, tst.expected, got)
			}
		}
	}
}
