package unique_elements

func UniqueElements(A []int) int {
	MergeSort(&A, 0, len(A)-1)
	d, c := 0, 0
	for i := 1; i < len(A); i++ {
		if A[i] == A[i-1] {
			d++
		} else if A[i]-A[i-1] == 1 {
			c += d
		} else if d > 0 {
			x := A[i] - A[i-1] - 1
			for ; x > 0 && d > 0; x-- {
				c += d
				d--
			}
			c += d
		}
	}
	c += (d * (d + 1)) / 2
	return c
}
func MergeSort(A *[]int, start, end int) {
	if start == end {
		return
	}
	mid := (start + end) / 2

	MergeSort(A, start, mid)
	MergeSort(A, mid+1, end)
	Merge(A, start, mid, end)
}
func Merge(A *[]int, start, mid, end int) {
	n, m := mid-start+1, end-mid
	b := make([]int, n)
	c := make([]int, m)

	for i := start; i <= mid; i++ {
		b[i-start] = (*A)[i]
	}

	for i := mid + 1; i <= end; i++ {
		c[i-mid-1] = (*A)[i]
	}
	i, j := 0, 0
	for k := start; k <= end; k++ {
		if i == n {
			(*A)[k] = c[j]
			j++
		} else if j == m {
			(*A)[k] = b[i]
			i++
		} else if b[i] <= c[j] {
			(*A)[k] = b[i]
			i++
		} else {
			(*A)[k] = c[j]
			j++
		}
	}
}
