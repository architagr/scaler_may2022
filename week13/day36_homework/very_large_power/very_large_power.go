package very_large_power

import "scaler-may-22/week13/day36_assingment/power_function"

func LargePower(A, B int) int {
	MOD := 1000000007
	A = A % MOD
	if B == 1 {
		return A
	}
	y := 1
	for i := 2; i <= B; i++ {
		y = (y * i) % (MOD - 1)
	}
	return power_function.Pow(A, y, MOD)
}
