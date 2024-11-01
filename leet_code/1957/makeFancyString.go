package _1957

func makeFancyString(s string) string {
	str := make([]byte, 0, len(s))
	str = append(str, s[0])

	for count, i := 1, 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			count = 1
		}
		if count < 3 {
			str = append(str, s[i])
		}
	}

	return string(str)
}
