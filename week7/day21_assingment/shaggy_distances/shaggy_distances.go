package shaggy_distances

func ShaggyDistances(A []int) int {
	mapA := make(map[int]int)
	n := len(A)
	min := n + 1
	for i := 0; i < n; i++ {
		if val, ok := mapA[A[i]]; ok {
			diff := i - val
			if diff < min {
				min = diff
			}
		} else {
			mapA[A[i]] = i
		}
	}
	if min == n+1 {
		return -1
	}
	return min
}
