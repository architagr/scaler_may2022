package another_count_rectangles

import (
	"fmt"
	"testing"
)

func TestSum2II(t *testing.T) {
	type TestCase struct {
		input    []int
		b        int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{1, 2},
		b:        5,
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 2},
		b:        1,
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := CountRectangles(testCase.input, testCase.b)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
