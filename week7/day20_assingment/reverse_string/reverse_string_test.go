package reverse_string

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputArr string
	expected string
}

func TestReverseString(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: "scaler",
		expected: "relacs",
	})

	testCases = append(testCases, TestCase{
		inputArr: "academy",
		expected: "ymedaca",
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := ReverseString(testcase.inputArr)
			if !reflect.DeepEqual(got, testcase.expected) {
				tb.Errorf("Tested %+v to  expected %+v but got %+v", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
