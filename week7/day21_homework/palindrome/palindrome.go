package palindrome

func CheckPalindrome(A string) int {
	charMap := createCharMap(A)
	oddCount := 0
	for _, v := range charMap {
		if v%2 == 1 {
			oddCount++
		}
	}
	if oddCount > 1 {
		return 0
	}
	return 1
}
func createCharMap(a string) map[byte]int {
	result := make(map[byte]int)
	n := len(a)
	for i := 0; i < n; i++ {
		val := result[a[i]]
		val++
		result[a[i]] = val
	}
	return result
}
