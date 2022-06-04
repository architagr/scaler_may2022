package row_to_column_zero

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input, expected [][]int
}

func TestRowToColumnZero(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3, 4}, {5, 6, 7, 0}, {9, 2, 0, 4}},
		expected: [][]int{{1, 2, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
	})

	testCases = append(testCases, TestCase{
		input: [][]int{
			{97, 18, 55, 1, 50, 17, 16, 0, 22, 14},
			{58, 14, 75, 54, 11, 23, 54, 95, 33, 23},
			{73, 11, 2, 80, 6, 43, 67, 82, 73, 4},
			{61, 22, 23, 68, 23, 73, 85, 91, 25, 7},
			{6, 83, 22, 81, 89, 85, 56, 43, 32, 89},
			{0, 6, 1, 69, 86, 7, 64, 5, 90, 37},
			{10, 3, 11, 33, 71, 86, 6, 56, 78, 31},
			{16, 36, 66, 90, 17, 55, 27, 26, 99, 59},
			{67, 18, 65, 68, 87, 3, 28, 52, 9, 70},
			{41, 19, 73, 5, 52, 96, 91, 10, 52, 21},
		},
		expected: [][]int{
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 14, 75, 54, 11, 23, 54, 0, 33, 23},
			{0, 11, 2, 80, 6, 43, 67, 0, 73, 4},
			{0, 22, 23, 68, 23, 73, 85, 0, 25, 7},
			{0, 83, 22, 81, 89, 85, 56, 0, 32, 89},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 3, 11, 33, 71, 86, 6, 0, 78, 31},
			{0, 36, 66, 90, 17, 55, 27, 0, 99, 59},
			{0, 18, 65, 68, 87, 3, 28, 0, 9, 70},
			{0, 19, 73, 5, 52, 96, 91, 0, 52, 21},
		},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case %d", (i+1)), func(tb *testing.T) {
			got := RowToColumnZero(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case no %d expected %+v, but got %+v", (i + 1), testCase.expected, testCase.input)
			}
		})
	}
}
