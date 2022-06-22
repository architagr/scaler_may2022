package toupper

func ToUpper(A []byte) []byte {
	n := len(A)
	for i := 0; i < n; i++ {
		if A[i] >= 'a' && A[i] <= 'z' {
			A[i] ^= 32
		}
	}

	return A
}
