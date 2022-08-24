package merge_two_sorted_arrays

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	a, b, expected []int
}

func TestMerge2SortedArray(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		a:        []int{4, 7, 9},
		b:        []int{2, 11, 19},
		expected: []int{2, 4, 7, 9, 11, 19},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Merge2SortedArray(testCase.a, testCase.b)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
