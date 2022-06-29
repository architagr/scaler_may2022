package check_palindome

func CheckPalindrome(A string) int {
	return validate(A, 0, len(A)-1)
}

func validate(A string, s, e int) int {
	if s > e {
		return 1
	}
	if A[s] != A[e] {
		return 0
	}
	return validate(A, s+1, e-1)
}
