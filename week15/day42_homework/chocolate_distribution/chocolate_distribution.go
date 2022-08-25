package chocolate_distribution

import (
	"math"
)

func ChocolateDistribution(A []int, B int) int {
	n := len(A)
	if B == 0 || n == 0 {
		return 0
	}
	// sort

	mergeSort(&A, 0, n-1)
	min := math.MaxInt32
	for i := 0; i <= n-B; i++ {
		if A[i+B-1]-A[i] < min {
			min = A[i+B-1] - A[i]
		}
	}
	return min
}

func mergeSort(A *[]int, start, end int) {
	if start == end {
		return
	}
	mid := (start + end) / 2
	mergeSort(A, start, mid)
	mergeSort(A, mid+1, end)
	merge(A, start, mid+1, end)
}
func merge(A *[]int, l, y, r int) {
	n := r - l + 1
	s := l
	x := y
	arr := make([]int, n)
	i := 0
	for ; l < x && y <= r && i < n; i++ {
		if (*A)[l] < (*A)[y] {
			arr[i] = (*A)[l]
			l++
		} else {
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
}
