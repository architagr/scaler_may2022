package coin_sum_infinite

func CoinSumInfinite(A []int, B int) int {
	n := len(A)
	MOD := 1000007
	dp := make([][]int, 2)
	dp[0] = make([]int, B+1)
	dp[1] = make([]int, B+1)

	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		dp[i%2][0] = 1
		for j := 1; j <= B; j++ {
			if j >= A[i-1] {
				dp[i%2][j] = (dp[(i-1)%2][j] + dp[i%2][j-A[i-1]]) % MOD
			} else {
				dp[i%2][j] = dp[(i-1)%2][j]
			}

		}
	}

	return dp[n%2][B]
}
