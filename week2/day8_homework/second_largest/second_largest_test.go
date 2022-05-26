package second_largest

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr []int
	expected int
}

func TestGetSecondHigest(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{2, 4, 3, 1, -1, 8, 5},
		expected: 5,
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{4, 4, 3, 1, -1, 8, 2},
		expected: 4,
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{1, 2, 3, 4, 5},
		expected: 4,
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := GetSecondHigest(testcase.inputArr)
			if got != testcase.expected {
				tb.Errorf("Tested %+v, expected %d but got %d", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
