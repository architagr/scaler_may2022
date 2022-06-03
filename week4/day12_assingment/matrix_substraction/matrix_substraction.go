package matrix_substraction

func MatrixSubstraction(A, B [][]int) [][]int {
	n := len(A)
	m := len(A[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			A[i][j] = A[i][j] - B[i][j]
		}
	}
	return A
}
