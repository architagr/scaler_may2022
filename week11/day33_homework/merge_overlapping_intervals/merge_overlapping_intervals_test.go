package merge_overlapping_intervals

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected []Interval
}

func TestMergeIntervals(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input: []Interval{
			{
				start: 1,
				end:   3,
			},
			{
				start: 2,
				end:   6,
			},
			{
				start: 8,
				end:   10,
			},
			{
				start: 15,
				end:   18,
			},
		},
		expected: []Interval{
			{
				start: 1,
				end:   6,
			},
			{
				start: 8,
				end:   10,
			},
			{
				start: 15,
				end:   18,
			},
		},
	})
	testCases = append(testCases, TestCase{
		input: []Interval{
			{
				start: 1,
				end:   3,
			},
			{
				start: 2,
				end:   6,
			},
			{
				start: 5,
				end:   10,
			},
			{
				start: 8,
				end:   18,
			},
		},
		expected: []Interval{
			{
				start: 1,
				end:   18,
			},
		},
	})

	testCases = append(testCases, TestCase{
		input: []Interval{
			{start: 5, end: 28},
			{start: 7, end: 70},
			{start: 54, end: 93},
			{start: 5, end: 98},
			{start: 46, end: 70},
			{start: 42, end: 63},
			{start: 5, end: 91},
			{start: 30, end: 49},
			{start: 36, end: 57},
			{start: 31, end: 95},
			{start: 86, end: 88},
			{start: 9, end: 90},
			{start: 5, end: 53},

			{start: 42, end: 62},
			{start: 14, end: 100},
			{start: 59, end: 75},
			{start: 32, end: 100},
			{start: 5, end: 79},
			{start: 31, end: 31},
			{start: 7, end: 42},
			{start: 13, end: 47},
			{start: 44, end: 87},
			{start: 61, end: 83},
			{start: 100, end: 100},
			{start: 96, end: 98},
			{start: 47, end: 51},

			{start: 34, end: 44},
			{start: 6, end: 53},
			{start: 30, end: 92},
			{start: 50, end: 64},
			{start: 37, end: 57},
			{start: 49, end: 67},
			{start: 2, end: 67},
			{start: 36, end: 50},
			{start: 55, end: 100},
			{start: 54, end: 78},
			{start: 58, end: 70},
			{start: 2, end: 37},
			{start: 13, end: 54},

			{start: 7, end: 60},
			{start: 16, end: 79},
			{start: 35, end: 78},
			{start: 17, end: 57},
			{start: 16, end: 84},
			{start: 60, end: 80},
			{start: 10, end: 54},
			{start: 54, end: 59},
			{start: 62, end: 85},
			{start: 7, end: 37},
			{start: 31, end: 99},
			{start: 40, end: 41},
			{start: 4, end: 99},

			{start: 28, end: 45},
			{start: 27, end: 71},
			{start: 14, end: 64},
		},
		expected: []Interval{
			{start: 2, end: 100},
		},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := MergeIntervals(testCase.input)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
