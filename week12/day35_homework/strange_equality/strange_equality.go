package strange_equality

func StrangeEquality(A int) int {
	ans := 0
	for i := 0; i < 32; i++ {
		x := 1 << i
		if x > A {
			ans += x
			break
		} else {
			if A&x == 0 {
				ans += x
			}
		}
	}

	return ans
}
