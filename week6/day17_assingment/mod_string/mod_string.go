package mod_string

import "strconv"

func findMod(A string, B int) int {
	var mod int64 = 0
	b := int64(B)
	mod10 := 10 % b
	last10Mod := 1 % b
	n := len(A)
	firstDigit, _ := strconv.Atoi(string(A[n-1]))
	getDigitMod := func(digit int) int64 {
		last10Mod = ((mod10 * last10Mod) % b)
		return ((int64(digit) % b) * last10Mod) % b
	}
	mod += (int64(firstDigit) % b * (last10Mod % b)) % b

	for i := n - 2; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(A[i]))
		mod += getDigitMod(digit)

	}
	return int(mod % b)
}
