package compare_sorted_subarrays

func CompareSortedSubarray(A []int, B [][]int) []int {
	m := len(B)
	hashMap := getHashMap(A)
	prefixSum := getPrefixSum(A, hashMap)

	arr := make([]int, len(B))

	for i := 0; i < m; i++ {
		l1, r1, l2, r2 := B[i][0], B[i][1], B[i][2], B[i][3]
		//Case1: when both subArrays are identical
		if l1 == l2 && r1 == r2 {
			arr[i] = 1
			continue
		}

		//Case2: when size of both the subArrays differs
		if r1-l1 != r2-l2 {
			arr[i] = 0
			continue
		}
		sum1, sum2 := prefixSum[r1], prefixSum[r2]
		if l1 > 0 {
			sum1 -= prefixSum[l1-1]
		}
		if l2 > 0 {
			sum2 -= prefixSum[l2-1]
		}
		if sum1 == sum2 {
			arr[i] = 1
		} else {
			arr[i] = 0
		}

	}
	return arr
}

func getPower(base, power, mod int) int {
	res := 1
	for power > 0 {
		if power%2 == 1 {
			res = ((res % mod) * (base % mod)) % mod
			power = power - 1
		}

		base = ((base % mod) * (base % mod)) % mod
		power = power / 2
	}

	return res
}

func getHashMap(A []int) map[int]int {
	hashMap := make(map[int]int)
	N := len(A)
	mod := 1000000007
	for i := 0; i < N; i++ {
		if _, ok := hashMap[A[i]]; !ok {
			val := getPower(A[i], A[i], mod)
			hashMap[A[i]] = val
		}
	}
	return hashMap
}

func getPrefixSum(A []int, hashMap map[int]int) []int64 {
	N := len(A)

	prefixSum := make([]int64, N)
	prefixSum[0] = int64(hashMap[A[0]])
	for i := 1; i < N; i++ {
		prefixSum[i] = (prefixSum[i-1] + int64(hashMap[A[i]]))
	}

	return prefixSum
}
