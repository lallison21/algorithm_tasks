package _2942

import (
	"fmt"
	"testing"
)

func TestFindWordsContaining(t *testing.T) {
	testCases := []struct {
		words    []string
		x        byte
		expected []int
	}{
		{
			words:    []string{"leet", "code"},
			x:        'e',
			expected: []int{0, 1},
		},
		{
			words:    []string{"abc", "bcd", "aaaa", "cbc"},
			x:        'a',
			expected: []int{0, 2},
		},
		{
			words:    []string{"abc", "bcd", "aaaa", "cbc"},
			x:        'z',
			expected: []int{},
		},
	}

	for index, testCase := range testCases {
		t.Run(fmt.Sprint(index), func(t *testing.T) {
			res := findWordsContaining(testCase.words, testCase.x)

			if len(res) != len(testCase.expected) {
				t.Fatalf("incorrect res len: expected %d, got: %d", len(testCase.expected), len(res))
			}

			for i := range len(res) {
				if res[i] != testCase.expected[i] {
					t.Fatalf("incorrect res: expected %d, got: %d", testCase.expected[i], res[i])
				}
			}
		})
	}
}
