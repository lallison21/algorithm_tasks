package _1957

import "testing"

func TestMakeFancyString(t *testing.T) {
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := makeFancyString(tt.input)
			if res != tt.want {
				t.Errorf("MakeFancyString() = %v, want %v", res, tt.want)
			}
		})
	}
}
