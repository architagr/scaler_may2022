package merge_intervals

type Interval struct {
	start, end int
}

func MergeInterval(intervals []Interval, newInterval Interval) []Interval {

	arr := make([]Interval, 0)

	for i := 0; i < len(intervals); i++ {
		if intervals[i].end < newInterval.start {
			arr = append(arr, intervals[i])
		} else if intervals[i].start > newInterval.end {
			arr = append(arr, newInterval)
			arr = append(arr, intervals[i:]...)
			return arr
		} else {
			newInterval.end = maxValue(newInterval.end, intervals[i].end)
			newInterval.start = minValue(newInterval.start, intervals[i].start)
		}
	}
	arr = append(arr, newInterval)
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
