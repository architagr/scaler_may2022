package minor_diagonal_sum

func MinorDiagonalSum(A [][]int) int {
	n := len(A)
	sum := A[0][n-1]

	for i := 1; i < n; i++ {
		sum += A[i][n-i-1]
	}
	return sum
}
