package find_fibonacci

func FindFibonacci(A int) int {
	if A < 2 {
		return A
	}
	a, b := 0, 1
	for i := 2; i <= A; i++ {
		a, b = b, a+b
	}
	return b
}
