package add_binary_strings

import "math"

func AddBinaryStrings(A, B string) string {
	ans, x := "", ""

	n, m := len(A), len(B)
	for i := 0; i < int(math.Abs(float64(n-m))); i++ {
		x += "0"
	}

	if m < n {
		B = x + B
	} else {
		A = x + A
	}
	n = len(A)
	carry := false
	for i := n - 1; i >= 0; i-- {
		if A[i] == '1' && B[i] == '1' && carry {
			ans = "1" + ans
			carry = true
		} else if A[i] == '1' && B[i] == '1' && !carry {
			ans = "0" + ans
			carry = true
		} else if (A[i] == '1' && B[i] == '0' && carry) || (A[i] == '0' && B[i] == '1' && carry) {
			ans = "0" + ans
			carry = true
		} else if (A[i] == '1' && B[i] == '0' && !carry) || (A[i] == '0' && B[i] == '1' && !carry) {
			ans = "1" + ans
			carry = false
		} else if A[i] == '0' && B[i] == '0' && !carry {
			ans = "0" + ans
			carry = false
		} else if A[i] == '0' && B[i] == '0' && carry {
			ans = "1" + ans
			carry = false
		}
	}
	if carry {
		ans = "1" + ans
	}
	return ans

}
