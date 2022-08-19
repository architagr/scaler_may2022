package prime_modulo_inverse

import "scaler-may-22/week13/day36_assingment/power_function"

func InverseModulo(A, B int) int {
	if A == 0 {
		return 0
	}
	return power_function.Pow(A, B-2, B)
}
