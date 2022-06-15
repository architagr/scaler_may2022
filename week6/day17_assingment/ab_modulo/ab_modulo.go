package ab_modulo

import "math"

func FindCommonNum(A, B int) int {
	// X = A - c*M --> 1
	// X = B - d*M --> 2
	// From 1 & 2
	// 0 = A-B - c*M + d*M
	// M = (A-B)/(d-c)
	// M will be max when d - c will be 1
	return int(math.Abs(float64(A - B)))
}
