package n_integers_containing_only_1_2_3

func GetNumbers(A int) []int {
	numbers := []int{1, 2, 3}
	if A <= 3 {
		return numbers[:A]
	}
	result := make([]int, A)
	result[0] = 1
	result[1] = 2
	result[2] = 3
	a := 0
	for x := 3; x < A; {
		result[x] = result[a]*10 + numbers[0]
		x++
		if x >= A {
			break
		}
		result[x] = result[a]*10 + numbers[1]
		x++
		if x >= A {
			break
		}
		result[x] = result[a]*10 + numbers[2]
		x++
		a++
	}
	return result
}
