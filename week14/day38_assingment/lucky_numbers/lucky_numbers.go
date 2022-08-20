package lucky_numbers

func LuckyNumber(A int) int {
	spf := spf(A)
	count := 0
	for i := 1; i <= A; i++ {
		set := make(map[int]bool)
		n := i
		for n > 1 {
			u := spf[n]
			for n%u == 0 {
				n /= u
			}
			set[u] = true
		}
		if len(set) == 2 {
			count++
		}
	}
	return count

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
