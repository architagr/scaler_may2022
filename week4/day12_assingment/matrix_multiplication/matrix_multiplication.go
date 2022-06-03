package matrix_multiplication

func FindMatrixMultiplication(A, B [][]int) [][]int {
	m := len(A)
	n := len(B)
	p := len(B[0])

	result := make([][]int, 0)

	for i := 0; i < m; i++ {
		sub := make([]int, 0)
		for j := 0; j < p; j++ {
			mul := 0
			for k := 0; k < n; k++ {
				mul += A[i][k] * B[k][j]
			}
			sub = append(sub, mul)
		}
		result = append(result, sub)
	}
	return result
}
