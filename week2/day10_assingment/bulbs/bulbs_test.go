package bulbs

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr []int
	expected int
}

func TestSolve(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputArr: []int{1, 0, 1, 0, 1, 0, 1, 0},
		expected: 7,
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{0, 1, 1, 0, 1, 0, 0, 1, 1},
		expected: 6,
	})
	testCases = append(testCases, TestCase{
		inputArr: []int{0, 1, 0, 1},
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{1, 1, 1, 1},
		expected: 0,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.inputArr), func(tb *testing.T) {
			got := Solve(testCase.inputArr)
			if got != testCase.expected {
				tb.Errorf("Tested %+v extected %d and got %d", testCase.inputArr, testCase.expected, got)
			}
		})
	}
}
