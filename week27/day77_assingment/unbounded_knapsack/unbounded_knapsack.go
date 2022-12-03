package unbounded_knapsack

func UnboundedKnapSack(cap int, values, weights []int) int {
	n := len(weights)
	dp := make([][]int, 2)
	dp[0] = make([]int, cap+1)
	dp[1] = make([]int, cap+1)

	for i := 1; i <= n; i++ {
		dp[i%2][0] = 0
		for j := 1; j <= cap; j++ {
			dp[i%2][j] = dp[(i-1)%2][j]
			if weights[i-1] <= j {
				x := dp[i%2][j-weights[i-1]] + values[i-1]
				dp[i%2][j] = maxValue(dp[i%2][j], x)
			}
		}
	}

	return dp[n%2][cap]
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
