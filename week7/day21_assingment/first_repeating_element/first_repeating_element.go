package first_repeating_element

func FindFirstRepeatingElement(A []int) int {
	n := len(A)
	intMap := make(map[int]int)
	minIndex := n + 1

	for i := 0; i < n; i++ {
		if val, ok := intMap[A[i]]; ok {
			if val < minIndex {
				minIndex = val
			}
		} else {
			intMap[A[i]] = i
		}
	}
	if minIndex == n+1 {
		return -1
	} else {
		return A[minIndex]
	}
}
