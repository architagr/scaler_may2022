package pair_sum_divisible_by_m

func PairSumDivisible(A []int, B int) int {
	MOD := 1000000007
	arr := make([]int, B)
	ans := 0

	for i := 0; i < len(A); i++ {
		arr[A[i]%B]++
	}
	ans = ((arr[0] * (arr[0] - 1)) / 2) % MOD

	for i := 1; i <= B/2; i++ {
		if B-i == i {
			ans += ((arr[i] * (arr[i] - 1)) / 2) % MOD
		} else {
			ans += (arr[i] % MOD * arr[B-i] % MOD) % MOD
		}
	}
	return ans % MOD
}
