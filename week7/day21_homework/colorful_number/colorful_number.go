package colorful_number

func CheckColorfulNumber(A int) int {
	a := splitDigits(A)
	prefixMul := getPrefixMul(a)
	m := make(map[int]int)
	n := len(a)
	flag := true
	for i := 0; i <= n; i++ {
		x := 1
		if i > 0 {
			x = prefixMul[i-1]
		}
		for j := i; j < n; j++ {
			mul := prefixMul[j] / x
			if _, ok := m[mul]; ok {
				flag = false
				break
			} else {
				m[mul] = mul
			}
		}
		if !flag {
			break
		}
	}
	if flag {
		return 1
	}
	return 0
}

func splitDigits(a int) []int {
	result := make([]int, 0)
	for a > 0 {
		result = append(result, a%10)
		a /= 10
	}
	return reverseArr(result)
}
func reverseArr(A []int) []int {
	i, j := 0, len(A)-1

	for i <= j {
		temp := A[i]
		A[i] = A[j]
		A[j] = temp
		i++
		j--
	}

	return A
}

func getPrefixMul(a []int) []int {
	mul := 1
	n := len(a)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		mul *= a[i]
		result[i] = mul
	}
	return result
}
