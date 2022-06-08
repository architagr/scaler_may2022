package christmas_trees

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, inputB []int
	expected      int
}

func TestFindMinCostCristmasTrees(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{5, 9, 10, 4, 7, 8},
		inputB:   []int{5, 6, 4, 7, 2, 5},
		expected: 12,
	})

	// testCases = append(testCases, TestCase{
	// 	input:    []int{1, 3, 5},
	// 	inputB:   []int{1, 2, 3},
	// 	expected: 6,
	// })
	// testCases = append(testCases, TestCase{
	// 	input:    []int{1, 6, 4, 2, 6, 9},
	// 	inputB:   []int{2, 5, 7, 3, 2, 7},
	// 	expected: 7,
	// })
	// testCases = append(testCases, TestCase{
	// 	input:    []int{2, 4, 5, 4, 10},
	// 	inputB:   []int{40, 30, 20, 10, 40},
	// 	expected: 90,
	// })
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case no %d", (i+1)), func(tb *testing.T) {
			got := FindMinCostCristmasTrees(testCase.input, testCase.inputB)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case number %d, expected %+v and got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
