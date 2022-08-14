package divisor_game

import "scaler-may-22/week13/day37_assingment/greatest_common_divisor"

func DivisorGame(A, B, C int) int {
	gcd := greatest_common_divisor.Gcd(B, C)
	lcm := (B * C) / gcd

	if lcm > A {
		return 0
	}
	return A / lcm
}
