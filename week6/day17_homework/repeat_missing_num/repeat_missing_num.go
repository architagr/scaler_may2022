package repeat_missing_num

import "math"

func FindReapeatingMissingNumber(A []int) []int {
	n := len(A)
	ans := make([]int, 0)
	for i := 0; i < n; i++ {

		idx := int(math.Abs(float64(A[i]))) - 1

		if A[idx] > 0 {
			A[idx] = -1 * A[idx]
		} else {
			ans = append(ans, idx+1)
		}

	}
	for i := 0; i < n; i++ {
		if A[i] > 0 {
			ans = append(ans, i+1)

		}
	}
	return ans
}
