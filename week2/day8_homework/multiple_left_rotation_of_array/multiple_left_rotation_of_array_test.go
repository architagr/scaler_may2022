package multiple_left_rotation_of_array

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

type TestCaseMultiple struct {
	inputArr     []int
	noOfRotation []int
	expected     [][]int
}

func TestMultipleLeftRotateArray(t *testing.T) {
	testCases := make([]TestCaseMultiple, 0)
	testCases = append(testCases, TestCaseMultiple{
		inputArr:     []int{1, 2, 3, 4, 5},
		noOfRotation: []int{2, 3},
		expected:     [][]int{{3, 4, 5, 1, 2}, {4, 5, 1, 2, 3}},
	})

	testCases = append(testCases, TestCaseMultiple{
		inputArr:     []int{5, 17, 100, 11},
		noOfRotation: []int{1},
		expected:     [][]int{{17, 100, 11, 5}},
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := MultipleLeftRotateArray(testcase.inputArr, testcase.noOfRotation)
			if !reflect.DeepEqual(got, testcase.expected) {
				tb.Errorf("Tested %+v to  expected %+v but got %+v", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
func TestLeftRotateArray(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr:     []int{1, 2, 3, 4, 5},
		noOfRotation: 3,
		expected:     []int{4, 5, 1, 2, 3},
	})

	testCases = append(testCases, TestCase{
		inputArr:     []int{12, 232, 35, 14, 5},
		noOfRotation: 7,
		expected:     []int{35, 14, 5, 12, 232},
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := LeftRotateArray(testcase.inputArr, testcase.noOfRotation)
			if !reflect.DeepEqual(got, testcase.expected) {
				tb.Errorf("Tested %+v to  expected %+v but got %+v", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
