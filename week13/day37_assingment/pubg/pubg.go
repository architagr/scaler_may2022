package pubg

import "scaler-may-22/week13/day37_assingment/greatest_common_divisor"

func Pubg(A []int) int {
	ans := A[0]
	for i := 1; i < len(A); i++ {
		ans = greatest_common_divisor.Gcd(ans, A[i])
	}
	return ans
}
