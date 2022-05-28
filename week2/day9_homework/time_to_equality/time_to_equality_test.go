package time_to_equality

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
		expected: 35,
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{2, 4, 1, 3, 2},
		expected: 8,
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %+v", testCase.inputArr), func(tb *testing.T) {
			got := TimeToEquality(testCase.inputArr)
			if got != testCase.expected {
				tb.Errorf("Tested %+v extected %d and got %d", testCase.inputArr, testCase.expected, got)
			}
		})
	}
}
