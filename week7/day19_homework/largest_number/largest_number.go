package largest_number

import (
	"fmt"
	"sort"
	"strconv"
)

func FindLargestNumber(A []int) string {

	sort.Slice(A, func(a, b int) bool {
		x, _ := strconv.Atoi(fmt.Sprintf("%d%d", A[a], A[b]))
		y, _ := strconv.Atoi(fmt.Sprintf("%d%d", A[b], A[a]))
		return x > y
	})
	str := ""
	if len(A) == 0 {
		str = "0"
	} else {
		zeroCount := 0
		for _, val := range A {
			if val == 0 {
				zeroCount++
			}
			str = fmt.Sprintf("%s%d", str, val)
		}
		if zeroCount == len(str) {
			str = "0"
		}

	}
	return str
}
