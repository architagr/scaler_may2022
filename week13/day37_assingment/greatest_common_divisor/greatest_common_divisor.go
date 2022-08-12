package greatest_common_divisor

func Gcd(A, B int) int {
	if A < B {
		A, B = B, A
	} else if A == B {
		return A
	}

	if B <= 0 && A > 1 {
		return A
	}
	if A <= 1 {
		return A
	}

	return Gcd(B, A%B)
}
