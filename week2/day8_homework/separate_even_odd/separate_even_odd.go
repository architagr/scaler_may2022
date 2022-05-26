package separate_even_odd

func SeparateEvenOdd(arr []int) (even []int, odd []int) {
	even = make([]int, 0)
	odd = make([]int, 0)

	for _, val := range arr {
		if val%2 == 0 {
			even = append(even, val)
		} else {
			odd = append(odd, val)
		}
	}
	return
}
