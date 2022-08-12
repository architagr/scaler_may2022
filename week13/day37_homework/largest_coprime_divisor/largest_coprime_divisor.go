package largest_coprime_divisor

import "scaler-may-22/week13/day37_assingment/greatest_common_divisor"

func CoPrimeDivisor(A, B int) int {
	gcd := greatest_common_divisor.Gcd(A, B)
	if gcd == 1 {
		return A
	}
	return CoPrimeDivisor(A/gcd, B)
}
