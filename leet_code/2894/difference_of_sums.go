package _2894

func differenceOfSums(n int, m int) int {
	num1, num2 := 0, 0

	for i := range n {
		num := i + 1

		switch num % m {
		case 0:
			num2 += num
		default:
			num1 += num
		}
	}

	return num1 - num2
}
