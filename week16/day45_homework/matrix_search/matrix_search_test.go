package matrix_search

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	b        int
	expected int
}

func TestSearchMatrix(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		b:        3,
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input: [][]int{
			{1, 3, 5, 7},
			{10, 11, 16, 20},
			{23, 30, 34, 50},
		},
		b:        4,
		expected: 0,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := SearchMatrix(testCase.input, testCase.b)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
