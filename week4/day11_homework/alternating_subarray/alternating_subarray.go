package alternating_subarray

func FindAlternatingSubArrayIndices(A []int, B int) []int {
	result := make([]int, 0)
	n := len(A)
	b := (2 * B) + 1
	for i := 0; i <= n-b; i++ {
		flag := true
		current := A[i]
		for j := 1; j < b; j++ {
			if A[i+j] == current {
				flag = false
				break
			} else {
				current = A[i+j]
			}
		}

		if flag {
			result = append(result, i+B)
		}
	}

	return result
}
