package maximum_sum_square_submatrix

import "scaler-may-22/common"

func MaxSqSubMatrix(A [][]int, B int) int {
	n := len(A)
	maxsum := int64(common.MinInt())
	prefixSum := findPrefixSum(A)
	for i := 0; i <= n-B; i++ {
		for j := 0; j <= n-B; j++ {
			topRow, leftCol := i, j
			buttomRow, rightCol := i+B-1, j+B-1

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
			maxsum = maxValue(maxsum, sum)
			if maxsum < sum {
				maxsum = sum
			}

		}
	}
	return int(maxsum)
}

func maxValue(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
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
