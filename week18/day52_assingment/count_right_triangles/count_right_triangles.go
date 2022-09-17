package count_right_triangles

func CountRightTriangles(A []int, B []int) int {
	MOD := int64(1000000007)
	ans := 0
	xPoints := make(map[int]int)
	yPoints := make(map[int]int)

	for i := 0; i < len(A); i++ {
		xPoints[A[i]]++
		yPoints[B[i]]++
	}

	for i := 0; i < len(A); i++ {
		x, y := A[i], B[i]

		countx, county := xPoints[x]-1, yPoints[y]-1
		ans = ans + int((int64(countx)%MOD*int64(county)%MOD)%MOD)
	}

	return ans
}
