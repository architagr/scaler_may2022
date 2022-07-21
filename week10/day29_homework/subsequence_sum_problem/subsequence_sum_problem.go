package subsequence_sum_problem

func SubsequenceSumProblem(A []int, B int) int {
	n := len(A)
	rowCount := 1 << n

	for i := 0; i < rowCount; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			x := 1 << j
			if i&x == x {
				sum += A[j]
			}
		}
		if sum == B {
			return 1
		}

	}
	return 0
}
