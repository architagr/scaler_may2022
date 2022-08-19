package compute_nCr_mod_p

import "scaler-may-22/week13/day36_assingment/power_function"

func ComputeNCR(A, B, C int) int {
	if A < B {
		return 0
	}
	// Base case
	if B == 0 {
		return 1
	}

	// Fill factorial array so that we
	// can find all factorial of r, n
	// and n-r
	fac := make([]int, A+1)
	fac[0] = 1

	for i := 1; i <= A; i++ {
		fac[i] = int((int64(fac[i-1]) * int64(i)) % int64(C))
	}
	x := modInverse(fac[B], C)
	y := modInverse(fac[A-B], C)
	return ((fac[A] * (x % C)) % C * (y % C)) % C
}

func modInverse(n, p int) int {
	return power_function.Pow(n, p-2, p)
}
