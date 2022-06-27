package common_elements

func FindCommonELements(A, B []int) []int {
	bMap := GetArrayCountMap(B)
	n := len(A)
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		val := bMap[A[i]]
		if val > 0 {
			val--
			bMap[A[i]] = val
			result = append(result, A[i])
		}
	}
	return result
}

func GetArrayCountMap(B []int) map[int]int {
	result := make(map[int]int)
	n := len(B)
	for i := 0; i < n; i++ {
		val := result[B[i]]
		val++
		result[B[i]] = val
	}

	return result
}
