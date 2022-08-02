package minimum_swaps_2

func MinSwaps(A []int, B int) int {
	countB := 0

	for i := 0; i < len(A); i++ {
		if A[i] <= B {
			countB++
		}
	}

	countInCurrWindow := 0 //count of el <= B in window (start, start+count)
	for i := 0; i < countB; i++ {
		if A[i] <= B {
			countInCurrWindow++
		}
	}

	minSwap := countB - countInCurrWindow //Minimum number of swaps required
	for i := 1; i < len(A)-countB+1; i++ {
		if A[i-1] <= B {
			countInCurrWindow--
		}
		if A[i+countB-1] <= B {
			countInCurrWindow++
		}
		if (countB - countInCurrWindow) < minSwap {
			minSwap = countB - countInCurrWindow
		}
	}

	return minSwap
}
