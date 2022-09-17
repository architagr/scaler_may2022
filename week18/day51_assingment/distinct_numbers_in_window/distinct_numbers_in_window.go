package distinct_numbers_in_window

func DistinctNumberInWindow(A []int, B int) []int {
	hashMap := make(map[int]int)
	n := len(A)
	result := make([]int, 0, n-B+1)
	if B > n {
		return result
	}
	for i := 0; i < B; i++ {
		hashMap[A[i]]++
	}
	result = append(result, len(hashMap))
	for i := B; i < n; i++ {
		hashMap[A[i-B]]--
		if hashMap[A[i-B]] == 0 {
			delete(hashMap, A[i-B])
		}
		hashMap[A[i]]++
		result = append(result, len(hashMap))

	}
	return result
}
