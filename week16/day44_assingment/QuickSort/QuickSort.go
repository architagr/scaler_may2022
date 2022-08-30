package QuickSort

func Sort(A []int) []int {
	QuickSort(&A, 0, len(A)-1)
	return A
}
func QuickSort(A *[]int, l, r int) {
	if l >= r {
		return
	}
	i, j := l+1, r
	for i <= j {
		if (*A)[i] <= (*A)[l] {
			i++
		} else if (*A)[j] > (*A)[l] {
			j--
		} else {
			(*A)[i], (*A)[j] = (*A)[j], (*A)[i]
			i++
			j--
		}
	}

	(*A)[i-1], (*A)[l] = (*A)[l], (*A)[i-1]

	QuickSort(A, l, i-2)
	QuickSort(A, i, r)
}
