package closest_palindrome

func CheckClosestPalindrome(A string) string {

	count := 0
	l, r := 0, len(A)-1
	for l <= r && count < 2 {
		if A[l] != A[r] {
			count++
		}
		l++
		r--

	}
	if count >= 2 || (len(A)&1 != 1 && count == 0) {
		return "NO"
	}
	return "YES"
}
