package matrix_multiplication

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, inputB, expected [][]int
}

func TestFindMatrixMultiplication(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   [][]int{{1, 2}, {3, 4}},
		inputB:   [][]int{{5, 6}, {7, 8}},
		expected: [][]int{{19, 22}, {43, 50}},
	})
	testCases = append(testCases, TestCase{
		inputA:   [][]int{{1, 1}},
		inputB:   [][]int{{2}, {3}},
		expected: [][]int{{5}},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindMatrixMultiplication(testCase.inputA, testCase.inputB)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
