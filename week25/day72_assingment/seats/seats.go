package seats

func seats(A string) int {
	MOD := 10000003
	ans := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 'x' {
			start := i
			end := i + 1
			for end < len(A) && A[end] != 'x' {
				end++
			}
			if end < len(A) {
				ans = (ans + (end - start - 1)) % MOD
			}
			i = end - 1

		}
	}

	return ans
}
