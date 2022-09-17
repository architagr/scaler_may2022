package longest_consecutive_sequence

func LongestConsecutiveSequence(A []int) int {
	hashMap := make(map[int]bool)
	ans := 0

	for i := 0; i < len(A); i++ {
		hashMap[A[i]] = true
	}
	for value := range hashMap {
		if _, ok := hashMap[value-1]; ok {
			continue
		}
		x := value
		y := 0
		_, ok := hashMap[x]
		for ok {
			y++
			x++
			_, ok = hashMap[x]
		}
		if y > ans {
			ans = y
		}
	}
	return ans
}
