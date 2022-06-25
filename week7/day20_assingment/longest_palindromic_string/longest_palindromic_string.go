package longest_palindromic_string

func FindLongestPalindrome(A string) string {

	max, n := 1, len(A)
	str := string(A[0])
	for i := 1; i < n-1; i++ {
		p1, p2 := i-1, i+1
		p11, p22 := checkPalandrome(A, p1, p2)
		if p11 < p1 && p22 > p2 {

			dif := p22 - p11 - 1
			if dif > max {
				max = dif
				str = getString(A, p11, p22)
			}

		}

	}

	for i := 0; i <= n-2; i++ {
		p1, p2 := i, i+1
		p11, p22 := checkPalandrome(A, p1, p2)
		if p11 < p1 && p22 > p2 {

			dif := p22 - p11 - 1
			if dif > max {
				max = dif
				str = getString(A, p11, p22)
			}

		}

	}

	return str
}

func checkPalandrome(A string, p1, p2 int) (int, int) {
	n := len(A)
	for p1 >= 0 && p2 < n && A[p1] == A[p2] {
		p1--
		p2++
	}
	return p1, p2
}

func getString(A string, p1, p2 int) string {
	p1++

	if p2 == len(A) {
		return A[p1:]
	}
	return A[p1:p2]
}
