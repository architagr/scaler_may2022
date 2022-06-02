package leader_in_array

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputArr []int
	expected []int
}

func TestFindLeaders(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: []int{16, 17, 4, 3, 5, 2},
		expected: []int{17, 2, 5},
	})

	testCases = append(testCases, TestCase{
		inputArr: []int{1, 2},
		expected: []int{2},
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := FindLeaders(testcase.inputArr)
			if !reflect.DeepEqual(got, testcase.expected) {
				tb.Errorf("Tested %+v to  expected %+v but got %+v", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
