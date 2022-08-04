package different_bits_sum_pairwise

func CountBits(A []int) int {
	ans := 0
	n := len(A) % 1000000007
	for i := 0; i < 32; i++ {
		countSet := 0
		for j := 0; j < len(A); j++ {
			if (A[j]>>i)&1 == 1 {
				countSet++
			}
		}
		countSet = countSet % 1000000007
		ans = (ans + (countSet * (n - countSet) % 1000000007)) % 1000000007
	}

	return (2 * ans) % 1000000007
}
