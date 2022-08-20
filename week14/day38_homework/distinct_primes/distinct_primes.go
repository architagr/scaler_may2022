package distinct_primes

import "math"

func DistinctPrimes(A []int) int {
	maxel := math.MinInt32
	for i := 0; i < len(A); i++ {
		maxel = maxValue(maxel, A[i])
	}
	spf := spf(maxel)
	set := make(map[int]bool)
	for i := 0; i < len(A); i++ {
		n := A[i]
		for n > 1 {
			u := spf[n]
			for n%u == 0 {
				n /= u
			}
			set[u] = true
		}
	}
	return len(set)
}

func spf(n int) []int {
	spf := make([]int, n+1)
	spf[0] = 1
	spf[1] = 1
	for i := 2; i <= n; i++ {
		spf[i] = i
	}
	for i := 2; i*i <= n; i++ {
		if spf[i] == i {
			for j := i * i; j <= n; j += i {
				spf[j] = minValue(spf[j], i)
			}
		}
	}
	return spf
}
func minValue(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
