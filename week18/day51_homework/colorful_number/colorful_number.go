package colorful_number

func ColorfulNumber(A int) int {
	arr := make([]int, 0, 40)
	hashMap := make(map[int]bool)

	for A > 0 {
		digit := A % 10
		if _, ok := hashMap[digit]; ok {
			return 0
		}
		hashMap[digit] = true
		arr = append(arr, digit)
		A /= 10
	}

	for i := 0; i < len(arr); i++ {
		m := arr[i]
		for j := i + 1; j < len(arr); j++ {
			m *= arr[j]
			if _, ok := hashMap[m]; ok {
				return 0
			}
			hashMap[m] = true
		}
	}

	return 1
}
