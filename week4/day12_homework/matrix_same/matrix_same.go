package matrix_same

func CompareMatrix(A, B [][]int) int {

	n := len(A)
	m := len(A[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if A[i][j] != B[i][j] {
				return 0
			}
		}
	}
	return 1
}
