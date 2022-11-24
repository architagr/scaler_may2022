package min_sum_path_in_matrix

import "math"

func MinPathSum(A [][]int) int {

	n, m := len(A), len(A[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			dp[i][j] = -1
		}
	}
	dp[0][0] = A[0][0]
	return waysToReach(n-1, m-1, A, dp)
}
func waysToReach(i, j int, A, dp [][]int) int {
	n, m := len(A), len(A[0])
	if i >= n || j >= m || i < 0 || j < 0 {
		return math.MaxInt
	}

	if dp[i][j] == -1 {
		dp[i][j] = minValue(waysToReach(i-1, j, A, dp), waysToReach(i, j-1, A, dp)) + A[i][j]
	}
	return dp[i][j]
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
