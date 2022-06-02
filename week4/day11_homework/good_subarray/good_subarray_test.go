package good_subarray

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr          []int
	inputNo, expected int
}

func TestFindCountOfGoodArray(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{1, 2, 3, 4, 5},
		inputNo:  4,
		expected: 6,
	})
	testCases = append(testCases, TestCase{
		inputArr: []int{13, 16, 16, 15, 9, 16, 2, 7, 6, 17, 3, 9},
		inputNo:  65,
		expected: 36,
	})
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing for %+v and number %d", testCase.inputArr, testCase.inputNo), func(tb *testing.T) {
			got := FindCountOfGoodArray(testCase.inputArr, testCase.inputNo)
			if got != testCase.expected {
				tb.Errorf("Tested %+v for %d, we expected %d and got %d", testCase.inputArr, testCase.inputNo, testCase.expected, got)
			}
		})
	}
}
