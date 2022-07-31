package sum_of_all_submatrices

func SumAllSubMatrix(A [][]int) int {
	sum := 0
	n, m := len(A), len(A[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			sum += (i + 1) * (j + 1) * (n - i) * (m - j) * A[i][j]
		}
	}
	return sum
}
