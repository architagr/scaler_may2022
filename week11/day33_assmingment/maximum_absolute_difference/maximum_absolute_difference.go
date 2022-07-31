package maximum_absolute_difference

import "scaler-may-22/common"

func MaxAbsoliteDiff(A []int) int {
	maxA := common.MinInt()
	minA := common.MaxInt()
	maxD := common.MinInt()
	minD := common.MaxInt()

	for i := 0; i < len(A); i++ {
		if A[i]+i < minA {
			minA = A[i] + i
		}
		if A[i]+i > maxA {
			maxA = A[i] + i
		}

		if A[i]-i < minD {
			minD = A[i] - i
		}
		if A[i]-i > maxD {
			maxD = A[i] - i
		}
	}
	if (maxA - minA) > (maxD - minD) {
		return (maxA - minA)
	}
	return (maxD - minD)
}
