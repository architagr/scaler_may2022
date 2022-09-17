package shaggy_and_distances

func MinDistanceSpecialPair(A []int) int {
	ans := len(A) + 1

	hashMap := make(map[int]int)
	for i := 0; i < len(A); i++ {
		if index, ok := hashMap[A[i]]; ok {
			if i-index < ans {
				ans = i - index
			}
		}
		hashMap[A[i]] = i
	}
	if ans == len(A)+1 {
		ans = -1
	}
	return ans
}
