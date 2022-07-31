package row_with_maximum_number_of_ones

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestMaxOnesRow(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    [][]int{{0, 1, 1}, {0, 0, 1}, {1, 1, 1}},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{0, 1, 1}, {0, 0, 1}, {0, 1, 1}},
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		input:    [][]int{{0, 0, 0, 0}, {0, 1, 1, 1}},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{0}},
		expected: -1,
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{1}},
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := MaxOnesRow(testCase.input)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
