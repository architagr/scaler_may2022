package armstrongnumbers

func GetArmstrongInRange(n int) []int {
	result := make([]int, 0)
	result = append(result, 1)
	for i := 2; i <= n; i++ {
		if checkArmstrong(i) {
			result = append(result, i)
		}
	}
	return result
}

func checkArmstrong(number int) bool {
	copy := number
	result := 0
	for copy != 0 {
		digit := copy % 10
		result += digit * digit * digit
		copy /= 10
	}
	return result == number
}
