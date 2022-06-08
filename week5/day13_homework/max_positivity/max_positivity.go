package max_positivity

func FindMaxPositivitySubArray(A []int) []int {
	result := make([]int, 0)
	temp := make([]int, 0)
	for _, val := range A {
		if val > -1 {
			temp = append(temp, val)
		} else {
			if len(temp) > len(result) {
				result = make([]int, 0)
				result = temp
			}
			temp = make([]int, 0)
		}
	}
	if len(temp) > len(result) {
		result = make([]int, 0)
		result = temp
	}
	return result
}
