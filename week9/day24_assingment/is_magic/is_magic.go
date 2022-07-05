package is_magic

func IsMagic(A int) int {
	for A > 9 {
		A = SumOfDigits(A)
	}
	if A == 1 {
		return 1
	}
	return 0
}

func SumOfDigits(A int) int {
	if A == 0 {
		return 0
	}
	return (A % 10) + SumOfDigits(A/10)
}
