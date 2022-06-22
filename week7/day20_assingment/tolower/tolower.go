package tolower

func ToLower(A []byte) []byte {
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i] >= 'A' && A[i] <= 'Z' {
			A[i] ^= 32
		}
	}

	return A
}
