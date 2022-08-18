package sum_of_digits

func SumOfDigits(A int) int {
	if A == 0 {
		return 0
	}
	return SumOfDigits(A/10) + (A % 10)
}
