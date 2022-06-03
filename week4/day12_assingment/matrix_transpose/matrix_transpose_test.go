package matrix_transpose

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected [][]int
}

func TestTransposeMatrix(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		expected: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2}, {1, 2}, {1, 2}},
		expected: [][]int{{1, 1, 1}, {2, 2, 2}},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case %d", (i+1)), func(tb *testing.T) {
			got := TransposeMatrix(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case no %d expected %+v, but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
