package merge_overlapping_intervals

import "sort"

type Interval struct {
	start, end int
}

func MergeIntervals(A []Interval) []Interval {
	sort.Slice(A, func(i, j int) bool {
		return A[i].start < A[j].start
	})
	arr := make([]Interval, 0, len(A))
	arr = append(arr, A[0])
	last := 0
	for i := 1; i < len(A); i++ {
		if A[i].start <= arr[last].end {
			arr[last].end = maxValue(arr[last].end, A[i].end)
			arr[last].start = minValue(arr[last].start, A[i].start)
		} else {
			arr = append(arr, A[i])
			last++
		}
	}
	return arr
}
func maxValue(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func minValue(a, b int) int {
	if a < b {
		return a
	}
	return b
}
