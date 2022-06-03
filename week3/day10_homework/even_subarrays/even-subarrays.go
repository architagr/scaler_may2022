package even_subarrays

func FindEvenSubarray(A []int) string {
	n := len(A)
	if n%2 != 0 || A[0]%2 != 0 || A[n-1]%2 != 0 {
		return "NO"
	} else {
		return "YES"
	}
}
