package element_removal

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestFindCostOfRemoval(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{3, 5, 1},
		expected: 14,
	})
	testCases = append(testCases, TestCase{
		input:    []int{2, 4, 1},
		expected: 11,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindCostOfRemoval(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
