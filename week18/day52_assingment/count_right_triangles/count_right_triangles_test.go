package count_right_triangles

import (
	"fmt"
	"testing"
)

func TestCountRightTriangles(t *testing.T) {
	type TestCase struct {
		inputA, inputB []int
		expected       int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 1, 2},
		inputB:   []int{1, 2, 1},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 1, 2, 3, 3},
		inputB:   []int{1, 2, 1, 2, 1},
		expected: 6,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := CountRightTriangles(testCase.inputA, testCase.inputB)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
