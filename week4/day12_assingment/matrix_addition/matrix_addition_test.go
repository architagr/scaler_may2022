package matrix_addition

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, inputB, expected [][]int
}

func TestMatrixAddition(t *testing.T) {
	testcases := make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		inputA:   [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		inputB:   [][]int{{9, 8, 7}, {6, 5, 4}, {3, 2, 1}},
		expected: [][]int{{10, 10, 10}, {10, 10, 10}, {10, 10, 10}},
	})

	testcases = append(testcases, TestCase{
		inputA:   [][]int{{1, 1}},
		inputB:   [][]int{{2, 3}},
		expected: [][]int{{3, 4}},
	})

	for i, testCase := range testcases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := MatrixAddition(testCase.inputA, testCase.inputB)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
