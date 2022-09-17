package count_rectangles

import (
	"fmt"
	"sort"
)

func CountRectangles(A, B []int) int {
	ans := 0
	points := make(map[string]int)
	taken := make(map[string]bool)
	for i := 0; i < len(A); i++ {
		points[fmt.Sprintf("%d_%d", A[i], B[i])] = i
	}
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A); j++ {
			if i == j {
				continue
			}
			xi, yi := A[i], B[i]
			xj, yj := A[j], B[j]
			p1, ok1 := points[fmt.Sprintf("%d_%d", xi, yj)]
			p2, ok2 := points[fmt.Sprintf("%d_%d", xj, yi)]

			if i != p1 && i != p2 && j != p1 && j != p2 {

				if ok1 && ok2 {
					a := []int{i, j, p1, p2}
					sort.Ints(a)
					s := fmt.Sprintf("%d", a)

					if _, ok := taken[s]; ok {
						continue
					}
					taken[s] = true
					ans++
				}
			}
		}
	}
	return ans
}
