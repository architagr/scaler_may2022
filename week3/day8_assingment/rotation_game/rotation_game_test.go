package rotationgame

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputArr     []int
	noOfRotation int
	expected     []int
}

func TestRotateArray(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr:     []int{1, 2, 3, 4, 5},
		noOfRotation: 1,
		expected:     []int{5, 1, 2, 3, 4},
	})

	testCases = append(testCases, TestCase{
		inputArr:     []int{12, 232, 35, 14, 5},
		noOfRotation: 7,
		expected:     []int{14, 5, 12, 232, 35},
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := RotateArray(testcase.inputArr, testcase.noOfRotation)
			if !reflect.DeepEqual(got, testcase.expected) {
				tb.Errorf("Tested %+v to  expected %+v but got %+v", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
