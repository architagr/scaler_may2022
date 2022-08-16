package prime_subsequence

import "math"

func CountPrimeSubsequence(A []int) int {
	count := 1
	n := len(A)
	for i := n - 1; i >= 0; i-- {
		if isPrime(A[i]) {
			count = (count * 2) % 1000000007
		}
	}

	return count - 1
}

func isPrime(a int) bool {
	if a == 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
		if a%i == 0 {
			return false
		}
	}
	return true
}
