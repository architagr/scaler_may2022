package find_subsequence

func SubSequence(A, B string) string {

	n, m := len(A), len(B)

	if n > m {
		return "NO"
	}
	j := 0
	for i := 0; i < m && j < n; i++ {
		if A[j] == B[i] {
			j++
		}
	}
	if j == n {
		return "YES"
	}
	return "NO"
}
