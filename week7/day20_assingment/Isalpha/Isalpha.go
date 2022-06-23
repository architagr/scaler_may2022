package Isalpha

func Isalpha(A []byte) int {

	for i := 0; i < len(A); i++ {
		if !((A[i] >= 'A' && A[i] <= 'Z') || (A[i] >= 'a' && A[i] <= 'z')) {
			return 0
		}
	}
	return 1
}
