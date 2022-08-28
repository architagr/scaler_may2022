package b_closest_points_to_origin

func BClosestPointToOrigin(A [][]int, B int) [][]int {
	if B == len(A) {
		return A
	}
	if B == 0 {
		return [][]int{}
	}
	MergeSort(&A, 0, len(A)-1)
	return A[:B]
}
func MergeSort(A *[][]int, start, end int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	MergeSort(A, start, mid)
	MergeSort(A, mid+1, end)
	Merge(A, start, mid, end)
}
func Merge(A *[][]int, start, mid, end int) {
	n, m := mid-start+1, end-mid
	b := make([][]int, n)
	c := make([][]int, m)

	for i := start; i <= mid; i++ {
		b[i-start] = make([]int, 2)
		b[i-start][0] = (*A)[i][0]
		b[i-start][1] = (*A)[i][1]
	}
	for i := mid + 1; i <= end; i++ {
		c[i-(mid+1)] = make([]int, 2)
		c[i-(mid+1)][0] = (*A)[i][0]
		c[i-(mid+1)][1] = (*A)[i][1]
	}
	i, j := 0, 0
	for k := start; k <= end; k++ {
		if i == n {
			(*A)[k][0] = c[j][0]
			(*A)[k][1] = c[j][1]
			j++
		} else if j == m {
			(*A)[k][0] = b[i][0]
			(*A)[k][1] = b[i][1]
			i++
		} else if ((b[i][0] * b[i][0]) + (b[i][1] * b[i][1])) <= ((c[j][0] * c[j][0]) + (c[j][1] * c[j][1])) {
			(*A)[k][0] = b[i][0]
			(*A)[k][1] = b[i][1]
			i++
		} else {
			(*A)[k][0] = c[j][0]
			(*A)[k][1] = c[j][1]
			j++
		}
	}
}
