package n_digit_numbers

import "fmt"

var dp map[string]int

func solve(A int, B int) int {
	dp = make(map[string]int)

	return count(A, B, 1)
}
func count(A, B, start int) int {
	if A == 0 && B == 0 {
		return 1
	}

	if B < 0 || A <= 0 {
		return 0
	}
	if val, ok := dp[fmt.Sprintf("%d_%d", A, B)]; ok {
		return val
	}
	var MOD = 1000000007
	ans := 0
	for i := start; i < 10; i++ {
		ans = (ans + count(A-1, B-i, 0)) % MOD
	}
	dp[fmt.Sprintf("%d_%d", A, B)] = ans
	return ans
}
