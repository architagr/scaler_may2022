package k_occurrences

func GetSumKOccurrences(A, B int, C []int) int {
	m := getCountMap(C)
	sum := 0
	flag := false
	for k, v := range m {
		if v == B {
			sum += k
			flag = true
		}
	}
	if !flag {
		return -1
	}
	return sum % 1000000007
}
func getCountMap(c []int) map[int]int {
	n := len(c)
	result := make(map[int]int)
	for i := 0; i < n; i++ {
		val := result[c[i]]
		result[c[i]] = val + 1
	}
	return result
}
