package sorted_permutation_rank

func FindRank(A string) int {
	MOD := 1000003

	arr1 := make([]int, 26)
	arr2 := make([]int, 26)
	mapChar := make(map[int]int)
	factArr := fact(len(A))
	for i := 0; i < len(A); i++ {
		if A[i] >= 'A' && A[i] <= 'Z' {
			arr1[A[i]-'A']++
		} else {
			arr2[A[i]-'a']++
		}
	}
	c := 0
	for i := 0; i < 26; i++ {
		if arr1[i] > 0 {
			mapChar['A'+i] = c
			c++
		}
	}
	for i := 0; i < 26; i++ {
		if arr2[i] > 0 {
			mapChar['a'+i] = c
			c++
		}
	}
	ans := 1
	length := len(A) - 1
	for i := 0; i < len(A); i, length = i+1, length-1 {
		countOfPrevChar := mapChar[int(A[i])]

		ans = ((ans)%MOD + (countOfPrevChar*factArr[length])%MOD) % MOD
		delete(mapChar, int(A[i]))
		for j := int(A[i]); j <= int('z'); j++ {
			if _, ok := mapChar[j]; ok {
				mapChar[j]--
			}
		}
	}
	return ans % MOD
}

func fact(a int) []int {
	arr := make([]int, a+1)
	arr[0] = 1
	arr[1] = 1
	if a <= 1 {
		return arr[:2]
	}
	ans := 1
	MOD := 1000003

	for i := 2; i <= a; i++ {

		ans = (ans * i) % MOD
		arr[i] = ans % MOD
	}
	return arr
}
