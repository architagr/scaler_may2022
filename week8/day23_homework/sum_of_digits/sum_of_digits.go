package sum_of_digits

func SumOfDigits(A int) int {
	if A == 0 {
		return 0
	}
	return (A % 10) + SumOfDigits(A/10)
}
