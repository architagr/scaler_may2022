package alternating_subarray

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputArr, expected []int
	inputNo            int
}

func TestFindAlternatingSubArrayIndices(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{1, 0, 1, 0, 1},
		inputNo:  1,
		expected: []int{1, 2, 3},
	})
	testCases = append(testCases, TestCase{
		inputArr: []int{0, 0, 0, 1, 1, 0, 1},
		inputNo:  0,
		expected: []int{0, 1, 2, 3, 4, 5, 6},
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{0, 0, 0, 1, 0, 0, 0, 1, 0, 1, 1},
		inputNo:  1,
		expected: []int{3, 7, 8},
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{1, 1, 1, 1, 1, 0, 1},
		inputNo:  0,
		expected: []int{0, 1, 2, 3, 4, 5, 6},
	})

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing for %+v and number %d", testCase.inputArr, testCase.inputNo), func(tb *testing.T) {
			got := FindAlternatingSubArrayIndices(testCase.inputArr, testCase.inputNo)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested %+v for %d, we expected %+v and got %+v", testCase.inputArr, testCase.inputNo, testCase.expected, got)
			}
		})
	}
}
