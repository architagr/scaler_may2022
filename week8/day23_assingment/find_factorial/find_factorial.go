package find_factorial

func findFactorial(A int) int {
	if A == 0 {
		return 1
	}
	return findFactorial(A-1) * A
}
