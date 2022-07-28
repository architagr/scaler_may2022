package max_non_negative_subArray

func MaxSet(A []int) []int {
	A = append(A, -1)
	i, start, sum, max, n := 0, -1, 0, 0, len(A)
	ans := make([]int, 0, n)
	for ; i < n; i++ {
		if A[i] >= 0 {
			start = i
			break
		}
	}

	for ; i < n; i++ {
		if A[i] < 0 {
			if sum > max || (sum == max && (i-start) > len(ans)) {
				ans = A[start:i]
				max = sum
			}
			sum = 0
			start = i + 1
		} else {
			sum += A[i]
		}
	}

	return ans
}
