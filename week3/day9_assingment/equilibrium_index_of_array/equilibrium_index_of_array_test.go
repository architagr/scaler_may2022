package equilibrium_index_of_array

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr []int
	expected int
}

func TestFindEquilibrimIndex(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{-7, 1, 5, 2, -4, 3, 0},
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{1, 2, 3},
		expected: -1,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.inputArr), func(tb *testing.T) {
			got := FindEquilibrimIndex(testCase.inputArr)
			if got != testCase.expected {
				tb.Errorf("Tested %+v extected %d and got %d", testCase.inputArr, testCase.expected, got)
			}
		})
	}
}
