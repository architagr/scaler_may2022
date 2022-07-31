package sub_matrix_sum_queries

func SubMatrixSumQueries(A [][]int, B, C, D, E []int) []int {
	prefixSum := findPrefixSum(A)
	result := make([]int, len(B))
	x := int64(1000000007)
	for i := 0; i < len(B); i++ {
		topRow, leftCol := B[i]-1, C[i]-1
		buttomRow, rightCol := D[i]-1, E[i]-1

		sum := int64(prefixSum[buttomRow][rightCol])
		if leftCol > 0 {
			sum -= int64(prefixSum[buttomRow][leftCol-1])
		}
		if topRow > 0 {
			sum -= int64(prefixSum[topRow-1][rightCol])
		}

		if leftCol > 0 && topRow > 0 {
			sum += int64(prefixSum[topRow-1][leftCol-1])
		}
		sum = sum % x
		if sum < 0 {
			sum += x
		}

		result[i] = int(sum)
	}
	return result
}

func findPrefixSum(A [][]int) [][]int {
	n, m := len(A), len(A[0])
	prefixSum := make([][]int, len(A))
	row := make([]int, len(A[0]))
	row[0] = A[0][0]
	for j := 1; j < m; j++ {
		row[j] = A[0][j] + row[j-1]
	}
	prefixSum[0] = row
	for i := 1; i < n; i++ {
		row = make([]int, len(A[0]))
		for j := 0; j < m; j++ {
			if j > 0 {
				row[j] = A[i][j] + row[j-1]
			} else {
				row[j] = A[i][j]
			}
		}
		prefixSum[i] = row
	}
	for i := 1; i < n; i++ {
		for j := 0; j < m; j++ {
			prefixSum[i][j] += prefixSum[i-1][j]
		}
	}
	return prefixSum
}
