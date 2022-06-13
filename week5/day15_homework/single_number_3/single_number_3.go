package single_number_3

func FundSingleNumer(A []int) []int {

	xorResponse := 0

	for _, val := range A {
		xorResponse ^= val
	}
	pos := findFirstSetBit(xorResponse)
	num := 1 << pos
	num1, num2 := 0, 0

	for _, val := range A {
		if val&num == num {
			num1 ^= val
		} else {
			num2 ^= val
		}
	}
	if num1 > num2 {
		return []int{num2, num1}
	} else {
		return []int{num1, num2}
	}
}

func findFirstSetBit(num int) int16 {
	var i, ans int16 = 0, 0

	for ; i < 32; i++ {
		x := (1 << i)
		if num&x == x {
			ans = i
			break
		}
	}
	return ans
}
