package _2942

import "strings"

func findWordsContaining(words []string, x byte) []int {
	res := make([]int, 0, len(words))

	for index, word := range words {
		if strings.ContainsAny(word, string(x)) {
			res = append(res, index)
		}
	}

	return res
}
