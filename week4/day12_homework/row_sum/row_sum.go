package row_sum

func FindRowSum(A [][]int) []int {
	result := make([]int, len(A))

	for i := 0; i < len(A); i++ {
		sum := 0

		for j := 0; j < len(A[i]); j++ {
			sum += A[i][j]
		}
		result[i] = sum
	}
	return result
}
