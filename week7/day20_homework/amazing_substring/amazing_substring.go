package amazing_substring

func FindAmazingSubstring(A string) int {
	count := int64(0)
	n := len(A)
	for i := 0; i < n; i++ {
		if checkVowel(A[i]) {
			count += int64(n - i)
		}
	}

	return int(count % 10003)
}

func checkVowel(a byte) bool {
	return a == 'a' || a == 'e' || a == 'i' || a == 'o' || a == 'u' || a == 'A' || a == 'E' || a == 'I' || a == 'O' || a == 'U'
}
