package check_palindrome_ii

func CheckPalindrome(A string) int {
	hash := make(map[byte]int)
	oddCount := 0

	for i := 0; i < len(A); i++ {
		hash[A[i]]++
	}
	for _, count := range hash {
		if count&1 == 1 {
			oddCount++
		}
	}
	if oddCount > 1 {
		return 0
	}
	return 1
}
