package divisibility_by_8

import "strconv"

func DivisibleBy8(A string) int {
	num := ""
	if len(A) >= 3 {
		num = A[len(A)-3:]
	} else {
		num = A
	}
	number, _ := strconv.Atoi(num)
	if number%8 == 0 {
		return 1
	}
	return 0
}
