package common_elements

func CommonElements(A, B []int) []int {
	n, m := len(A), len(B)
	arr := make([]int, 0, n+m)
	hashA, hashB := make(map[int]int), make(map[int]int)

	for i := 0; i < n; i++ {
		hashA[A[i]]++
	}

	for i := 0; i < m; i++ {
		hashB[B[i]]++
	}

	for key, countA := range hashA {
		if countB, ok := hashB[key]; ok {
			count := minValue(countA, countB)
			for i := 0; i < count; i++ {
				arr = append(arr, key)
			}
		}
	}
	return arr
}

func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
