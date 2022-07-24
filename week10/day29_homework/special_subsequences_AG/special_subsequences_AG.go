package special_subsequences_AG

func SpecialSubsequencesAG(A string) int {
	gcount, n, sum := 0, len(A), 0

	for i := n - 1; i >= 0; i-- {
		if A[i] == 'A' {
			sum += gcount % 1000000007
		}
		if A[i] == 'G' {
			gcount++
		}
	}
	return sum % 1000000007
}
