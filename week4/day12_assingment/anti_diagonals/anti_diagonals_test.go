package anti_diagonals

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected [][]int
}

func TestGetAntiDiagonal(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		expected: [][]int{{1, 0, 0}, {2, 4, 0}, {3, 5, 7}, {6, 8, 0}, {9, 0, 0}},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case %d", (i+1)), func(tb *testing.T) {
			got := GetAntiDiagonal(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case no %d expected %+v, but got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
