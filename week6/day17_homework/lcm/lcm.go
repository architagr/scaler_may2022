package lcm

func FindLcm(A, B int) int {
	hcf := FindHcf(A, B)

	return (A * B) / hcf
}

func FindHcf(A, B int) int {
	hcf := 1
	min := min(A, B)
	for i := min; i > 1; i-- {
		if A%i == 0 && B%i == 0 {
			hcf = i
			break
		}
	}
	return hcf
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
