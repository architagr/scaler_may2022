package matrix_transpose

func TransposeMatrix(A [][]int) [][]int {
	n := len(A)
	m := len(A[0])
	result := make([][]int, m)
	for j := 0; j < m; j++ {
		sub := make([]int, n)
		for i := 0; i < n; i++ {
			sub[i] = A[i][j]
		}
		result[j] = sub
	}
	return result
}
