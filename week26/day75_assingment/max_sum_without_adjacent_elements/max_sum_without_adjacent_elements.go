package max_sum_without_adjacent_elements

func adjacent(A [][]int) int {
	dp := make([][]int, 2)
	m := len(A[0])
	dp[0] = make([]int, m)
	dp[1] = make([]int, m)
	for j := 0; j < m; j++ {
		dp[0][j] = -1
		dp[1][j] = -1
	}
	j := m - 1
	max := 0
	for k := 0; k < 2 && j >= 0; k, j = k+1, j-1 {
		dp[0][j] = A[0][j]
		if max < dp[0][j] {
			max = dp[0][j]
		}
		dp[1][j] = A[1][j]
		if max < dp[1][j] {
			max = dp[1][j]
		}
	}

	maxa := 0
	for ; j >= 0; j-- {
		dp[0][j], maxa = calc(0, j, maxa, A, dp)
		dp[1][j], maxa = calc(1, j, maxa, A, dp)
		if maxa > max {
			max = maxa
		}
	}
	for j := 0; j < 2 && j < m; j++ {
		if max < dp[1][j] {
			max = dp[1][j]
		}
		if max < dp[0][j] {
			max = dp[0][j]
		}
	}
	return max
}
func calc(i, j, max int, A, dp [][]int) (int, int) {

	if max < dp[0][j+2] {
		max = dp[0][j+2]
	}
	if max < dp[1][j+2] {
		max = dp[1][j+2]
	}

	return A[i][j] + max, max
}
