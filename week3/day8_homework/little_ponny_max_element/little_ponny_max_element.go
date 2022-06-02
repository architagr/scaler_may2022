package little_ponny_max_element

func LittlePonnyMaxElement(A []int, B int) int {
	bCheck := false
	count := 0

	for _, val := range A {
		if val == B {
			bCheck = true
		}
		if val > B {
			count++
		}
	}

	if bCheck {
		return count
	}
	return -1
}
