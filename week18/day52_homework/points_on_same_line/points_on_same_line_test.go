package points_on_same_line

import (
	"fmt"
	"testing"
)

func TestPointsOnSameLine(t *testing.T) {
	type TestCase struct {
		inputA, inputB []int
		expected       int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{-1, 0, 1, 2, 3, 3},
		inputB:   []int{1, 0, 1, 2, 3, 4},
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{3, 1, 4, 5, 7, -9, -8, 6},
		inputB:   []int{4, -8, -3, -2, -1, 5, 7, -4},
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{2, 1, -7, 4, 1, -2, -7, 5},
		inputB:   []int{-6, -7, 3, -10, 7, -10, 1, 2},
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := PointsOnSameLine(testCase.inputA, testCase.inputB)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
