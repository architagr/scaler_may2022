package goodpair

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr []int
	inputB   int
	expected int
}

func TestCheckGoodPairExists(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{1, 2, 3, 4, 5},
		inputB:   9,
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{12, 232, 35, 14, 5},
		inputB:   10,
		expected: 0,
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := CheckGoodPairExists(testcase.inputArr, testcase.inputB)
			if got != testcase.expected {
				tb.Errorf("Tested %+v to have %d expected %d but got %d", testcase.inputArr, testcase.inputB, testcase.expected, got)
			}
		})
	}
}
