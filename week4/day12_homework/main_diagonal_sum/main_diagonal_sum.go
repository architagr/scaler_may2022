package main_diagonal_sum

func MainDiagonalSum(A [][]int) int {
	sum := A[0][0]

	for i := 1; i < len(A); i++ {
		sum += A[i][i]
	}
	return sum
}
