package first_repeating_element

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestFindFirstRepeatingElement(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{10, 5, 3, 4, 3, 5, 6},
		expected: 5,
	})
	testCases = append(testCases, TestCase{
		input:    []int{6, 10, 5, 4, 9, 120},
		expected: -1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindFirstRepeatingElement(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
