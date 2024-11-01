package _1957

import "testing"

func BenchmarkMakeFancyString(b *testing.B) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "s = leeetcode",
			input: "leeetcode",
			want:  "leetcode",
		},
		{
			name:  "s = aaabaaaa",
			input: "aaabaaaa",
			want:  "aabaa",
		},
		{
			name:  "s = aab",
			input: "aab",
			want:  "aab",
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			res := makeFancyString(test.input)
			if res != test.want {
				b.Fatal("expected", test.want, "got", res)
			}
		}
	}
}
