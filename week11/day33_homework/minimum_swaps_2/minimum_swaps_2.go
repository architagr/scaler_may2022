package minimum_swaps_2

func MinSwaps2(A []int) int {
	ans := 0

	for i := 0; i < len(A); {
		if A[i] != i {
			A[i], A[A[i]] = A[A[i]], A[i]
			ans++
			continue
		}
		i++
	}
	return ans
}
