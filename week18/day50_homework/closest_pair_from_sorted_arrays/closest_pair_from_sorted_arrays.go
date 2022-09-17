package closest_pair_from_sorted_arrays

import "math"

func FindClosestPair(A, B []int, C int) []int {
	result := make([]int, 2)

	n, m := len(A), len(B)
	i, j := 0, m-1
	mini, minj := A[n-1], B[m-1]
	diff := float64(math.MaxInt)
	for i < n && j >= 0 {
		sum := B[j] + A[i]
		d := math.Abs(float64(sum - C))
		if d < diff || (d == diff && A[i] <= mini && B[j] <= minj) {
			result[0] = A[i]
			result[1] = B[j]
			diff = d
			mini = A[i]
			minj = B[j]
		}

		if sum < C {
			i++
		} else if sum > C {
			j--
		} else {
			return result
		}
	}
	return result
}
