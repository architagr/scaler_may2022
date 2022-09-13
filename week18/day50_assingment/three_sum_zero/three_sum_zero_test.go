package three_sum_zero

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	type TestCase struct {
		input    []int
		expected [][]int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{-1, 0, 1, 2, -1, 4},
		expected: [][]int{{-1, 0, 1}, {-1, -1, 2}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ThreeSum(testCase.input)
			if !reflect.DeepEqual(testCase.expected, got) {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
