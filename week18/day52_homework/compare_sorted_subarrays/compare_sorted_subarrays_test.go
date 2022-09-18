package compare_sorted_subarrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompareSortedSubarray(t *testing.T) {
	type TestCase struct {
		inputA   []int
		inputB   [][]int
		expected []int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA: []int{100000, 100000, 100000, 100000, 100000},
		inputB: [][]int{{0, 3, 1, 4},
			{0, 1, 2, 3},
			{4, 4, 1, 1},
			{1, 3, 0, 0},
			{2, 4, 1, 1}},
		expected: []int{1, 1, 1, 0, 0},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := CompareSortedSubarray(testCase.inputA, testCase.inputB)
			if !reflect.DeepEqual(testCase.expected, got) {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
