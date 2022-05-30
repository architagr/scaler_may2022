package bulbs

func Solve(A []int) int {
	count := 0
	if len(A) == 0 {
		return 0
	}

	if A[0] == 0 {
		count = 1
	}
	for i := 1; i < len(A); i++ {
		if A[i] != A[i-1] {
			count++
		}
	}
	return count
}
