package balanced_paranthesis

func BalancedParanthesis(A string) int {
	n := len(A)
	stack := make([]byte, n)
	lastElement := -1
	m := map[byte]byte{'(': ')'}

	for i := 0; i < n; i++ {
		if lastElement < 0 {
			lastElement++
			stack[lastElement] = A[i]
			continue
		}

		v := m[stack[lastElement]]
		if v == A[i] {
			lastElement--
		} else {
			lastElement++
			stack[lastElement] = A[i]
		}

	}
	if lastElement != -1 {
		return 0
	}
	return 1
}
