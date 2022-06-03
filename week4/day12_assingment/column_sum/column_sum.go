package column_sum

func MatrixColumnSum(A [][]int) []int {
	n := len(A)
	m := len(A[0])
	result := make([]int, m)

	for j := 0; j < m; j++ {
		sum := 0
		for i := 0; i < n; i++ {
			sum += A[i][j]
		}
		result[j] = sum
	}

	return result
}
