package pick_from_both_sides

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr []int
	inputB   int
	expected int
}

func TestPickFromBothSides(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{-7, 1, 5, 2, -4, 3, 0},
		inputB:   2,
		expected: 3,
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{1, 2, 3},
		inputB:   1,
		expected: 3,
	})
	testCases = append(testCases, TestCase{
		inputArr: []int{-7, 1, 5, 2, -4, 3, 0},
		inputB:   7,
		expected: 0,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v for %d", testCase.inputArr, testCase.inputB), func(tb *testing.T) {
			got := PickFromBothSides(testCase.inputArr, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("Tested %+v for %d extected %d and got %d", testCase.inputArr, testCase.inputB, testCase.expected, got)
			}
		})
	}
}
