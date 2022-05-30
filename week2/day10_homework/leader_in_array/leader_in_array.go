package leader_in_array

func FindLeaders(A []int) []int {
	n := len(A)
	result := make([]int, 0, n)
	if n == 0 {
		return result
	}
	result = append(result, A[n-1])
	max := A[n-1]
	for i := n - 1; i >= 0; i-- {
		if A[i] > max {
			result = append(result, A[i])
			max = A[i]
		}
	}
	return result
}
