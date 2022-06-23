package count_occurrences

func CountOccurrence(A string) int {
	count := 0
	n := len(A)
	for i := 0; i <= n-3; i++ {
		if A[i] == 'b' && A[i+1] == 'o' && A[i+2] == 'b' {
			count++
		}
	}
	return count
}
