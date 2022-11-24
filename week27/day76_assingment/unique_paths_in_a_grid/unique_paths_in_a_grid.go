package unique_paths_in_a_grid

func UniquePathsWithObstacles(A [][]int) int {
	n, m := len(A), len(A[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if A[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = -1
			}
		}
	}
	if dp[0][0] == 0 {
		return 0
	}
	dp[0][0] = 1
	return waysToReach(n-1, m-1, A, dp)
}
func waysToReach(i, j int, A, dp [][]int) int {
	n, m := len(A), len(A[0])
	if i >= n || j >= m || i < 0 || j < 0 {
		return 0
	}
	if dp[i][j] == 0 {
		return 0
	}
	if dp[i][j] == -1 {
		dp[i][j] = waysToReach(i-1, j, A, dp) + waysToReach(i, j-1, A, dp)
	}
	return dp[i][j]
}
