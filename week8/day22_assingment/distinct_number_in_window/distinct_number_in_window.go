package distinct_number_in_window

func DistinctNumbers(A []int, B int) []int {

	mapA := make(map[int]int)
	result := make([]int, 0)
	n := len(A)
	for i := 0; i < B; i++ {
		val := mapA[A[i]]
		val++
		mapA[A[i]] = val
	}
	result = append(result, len(mapA))
	for i := B; i < n; i++ {
		val := mapA[A[i-B]]
		val--
		mapA[A[i-B]] = val
		if val == 0 {
			delete(mapA, A[i-B])
		}

		x := mapA[A[i]]
		mapA[A[i]] = x + 1

		result = append(result, len(mapA))
	}
	return result
}
