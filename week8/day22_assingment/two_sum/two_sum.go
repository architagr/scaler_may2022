package two_sum

func TwoSum(A []int, B int) []int {
	mapA := make(map[int]int)
	n := len(A)

	for i := 0; i < n; i++ {
		if index, ok := mapA[B-A[i]]; ok {
			return []int{index + 1, i + 1}
		}
		if _, ok := mapA[A[i]]; !ok {
			mapA[A[i]] = i
		}
	}
	return []int{}
}
