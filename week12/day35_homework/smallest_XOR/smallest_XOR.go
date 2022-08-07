package smallest_XOR

func SmallestXor(A, B int) int {
	count := 0
	arr := make([]bool, 32)
	ans := 0
	for i := 0; i < 32; i++ {
		if (1<<i)&A != 0 {
			count++
			arr[i] = true
		}
	}
	if count > B {
		for i := 31; B > 0 && i >= 0; i-- {
			if arr[i] {
				ans += 1 << i
				B--
			}
		}
	} else if count < B {
		B -= count
		ans = A
		for i := 0; B > 0 && i < 32; i++ {
			if !arr[i] {
				ans += 1 << i
				B--
			}
		}
	} else {
		ans = A
	}
	return ans
}
