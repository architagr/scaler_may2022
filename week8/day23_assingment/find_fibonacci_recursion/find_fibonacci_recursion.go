package find_fibonacci_recursion

func findAthFibonacci(A int) int {
	if A <= 1 {
		return A
	}
	return findAthFibonacci(A-1) + findAthFibonacci(A-2)
}
