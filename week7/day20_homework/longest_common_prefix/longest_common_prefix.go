package longest_common_prefix

func FindLongestCommonPrefix(A []string) string {
	minString := findMinLengthString(A)
	n, nm := len(A), len(minString)
	flag := true
	j := 0
	for ; j < nm; j++ {
		for i := 0; i < n; i++ {
			if minString[j] != A[i][j] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
	}

	return getStringPart(minString, 0, j)
}
func getStringPart(A string, start, end int) string {
	if end == len(A)-1 {
		return A[start:]
	}
	return A[start:end]

}
func findMinLengthString(A []string) string {
	n := len(A)
	min := 1000001
	str := ""
	for i := 0; i < n; i++ {
		if min > len(A[i]) {
			min = len(A[i])
			str = A[i]
		}
	}
	return str
}
