package min_picks

import "log"

func GetMinPicks(A []int) int {

	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	min, max := MaxInt, MinInt

	for _, val := range A {
		if val%2 == 0 && val > max {
			max = val
		}
		if val%2 != 0 && val < min {
			min = val
		}
	}
	log.Printf("max %d, min %d", max, min)
	return max - min
}
