package _2894

import (
	"fmt"
	"testing"
)

func TestDifferenceOfSums(t *testing.T) {
	testCases := []struct {
		n        int
		m        int
		expected int
	}{
		{
			n:        10,
			m:        3,
			expected: 19,
		},
		{
			n:        5,
			m:        6,
			expected: 15,
		},
		{
			n:        5,
			m:        1,
			expected: -15,
		},
	}

	for index, tc := range testCases {
		t.Run(fmt.Sprintf("%d", index), func(t *testing.T) {
			result := differenceOfSums(tc.n, tc.m)
			if result != tc.expected {
				t.Errorf("expected: %d, got: %d", tc.expected, result)
			}
		})
	}
}
