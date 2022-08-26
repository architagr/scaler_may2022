package inversion_count_in_an_array

func GetInversionCount(A []int) int {
	return mergeSort(&A, 0, len(A)-1)
}

func mergeSort(A *[]int, start, end int) int {
	if start == end {
		return 0
	}
	mid := (start + end) / 2
	a := mergeSort(A, start, mid)
	b := mergeSort(A, mid+1, end)
	ab := merge(A, start, mid+1, end)
	return (a + b + ab) % (1000000007)
}
func merge(A *[]int, l, y, r int) int {
	n := r - l + 1
	s := l
	x := y
	arr := make([]int, n)
	i := 0
	c := 0
	for ; l < x && y <= r && i < n; i++ {
		if (*A)[l] <= (*A)[y] {
			arr[i] = (*A)[l]
			l++
		} else {
			c = (c + (x - l)) % (1000000007)
			arr[i] = (*A)[y]
			y++
		}
	}

	if l == x && y <= r {
		for y <= r {
			arr[i] = (*A)[y]
			y++
			i++
		}
	}
	if y > r && l < x {
		for l < x {
			arr[i] = (*A)[l]
			l++
			i++
		}
	}

	i = 0
	for ; i < n; i++ {
		(*A)[s+i] = arr[i]
	}
	return c % (1000000007)
}
