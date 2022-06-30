package largest_number

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected string
}

func TestFindLargestNumber(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []int{3, 30, 34, 5, 9},
		expected: "9534330",
	})
	testCases = append(testCases, TestCase{
		input:    []int{2, 3, 9, 0},
		expected: "9320",
	})
	testCases = append(testCases, TestCase{
		input:    []int{0, 0, 0, 0},
		expected: "0",
	})
	testCases = append(testCases, TestCase{
		input:    []int{472, 663, 964, 722, 485, 852, 635, 4, 368, 676, 319, 412},
		expected: "9648527226766636354854724412368319",
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindLargestNumber(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %s but got %s", (i + 1), testCase.expected, got)
			}
		})
	}

}
