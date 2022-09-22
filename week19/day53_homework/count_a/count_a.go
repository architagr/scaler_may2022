package count_a

func CountA(A string) int {
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 'a' {
			count++
		}
	}
	return (count * (count + 1)) / 2
}
