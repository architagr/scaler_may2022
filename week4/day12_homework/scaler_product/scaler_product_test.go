package scaler_product

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    [][]int
	inputB   int
	expected [][]int
}

func TestFindScalerProduct(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 2, 3, 4}},
		inputB:   2,
		expected: [][]int{{2, 4, 6, 8}, {10, 12, 14, 16}, {18, 4, 6, 8}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case no %d", (i+1)), func(tb *testing.T) {
			got := FindScalerProduct(testCase.input, testCase.inputB)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case number %d, expected %+v and got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
