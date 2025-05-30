package _2131

import (
	"fmt"
	"testing"
)

func TestLongestPalondrome(t *testing.T) {
	testCases := []struct {
		input    []string
		expected int
	}{
		{
			input:    []string{"lc", "cl", "gg"},
			expected: 6,
		},
		{
			input:    []string{"ab", "ty", "yt", "lc", "cl", "ab"},
			expected: 8,
		},
		{
			input:    []string{"cc", "ll", "xx"},
			expected: 2,
		},
	}

	for index, tc := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			res := longestPalindrome(tc.input)
			if res != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, res)
			}
		})
	}
}
