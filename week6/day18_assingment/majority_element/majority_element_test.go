package majority_element

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestFindMajorityElement(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{2, 1, 2},
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		input:    []int{1, 3, 4, 3, 2, 3, 3, 3, 3, 2, 2},
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		input:    []int{1, 3, 4, 3, 2, 3, 3, 1, 8, 2, 2},
		expected: -1,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", testCase.input), func(tb *testing.T) {
			got := findMajorityElement(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}

}
