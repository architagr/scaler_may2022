package reversearray

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputArr []int
	expected []int
}

func TestReverseArray(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{1, 2, 3, 4, 5},
		expected: []int{5, 4, 3, 2, 1},
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{12, 232, 35, 14, 5},
		expected: []int{5, 14, 35, 232, 12},
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := ReverseArray(testcase.inputArr, 0, len(testcase.inputArr)-1)
			if !reflect.DeepEqual(got, testcase.expected) {
				tb.Errorf("Tested %+v to  expected %+v but got %+v", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
