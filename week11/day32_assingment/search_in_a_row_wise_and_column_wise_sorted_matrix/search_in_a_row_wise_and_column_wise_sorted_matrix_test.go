package search_in_a_row_wise_and_column_wise_sorted_matrix

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input       [][]int
	b, expected int
}

func TestSearch(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		b:        2,
		expected: 1011,
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		b:        20,
		expected: -1,
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{2, 8, 8, 8}, {2, 8, 8, 8}, {2, 8, 8, 8}},
		b:        8,
		expected: 1011,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d ", (i+1)), func(tb *testing.T) {
			got := Search(testCase.input, testCase.b)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
