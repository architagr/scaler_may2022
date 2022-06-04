package rotate_matrix

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected [][]int
}

func TestRotateMatrix(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
		expected: [][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case %d", (i+1)), func(tb *testing.T) {
			RotateMatrix(&testCase.input)
			if !reflect.DeepEqual(testCase.input, testCase.expected) {
				tb.Errorf("tested case no %d expected %+v, but got %+v", (i + 1), testCase.expected, testCase.input)
			}
		})
	}
}
