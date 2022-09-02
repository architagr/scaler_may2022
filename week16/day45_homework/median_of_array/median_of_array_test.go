package median_of_array

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB []int
	expected       float64
}

func TestFindMedianSortedArrays(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 4, 5},
		inputB:   []int{2, 3},
		expected: 3.0,
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{1, 2, 3},
		inputB:   []int{4},
		expected: 2.5,
	})
	testCases = append(testCases, TestCase{
		inputA:   []int{4},
		inputB:   []int{1, 2, 3},
		expected: 2.5,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{-35, -30, 22, 27, 33, 35, 40, 50},
		inputB:   []int{-39, -28, -25, -17, 1, 30, 33, 37},
		expected: 24.5,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindMedianSortedArrays(testCase.inputA, testCase.inputB)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %f but we got %f", (i + 1), testCase.expected, got)
			}
		})
	}
}
