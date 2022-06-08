package add_binary

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, inputB, expected string
}

func TestAddBinary(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   "100",
		inputB:   "11",
		expected: "111",
	})
	testCases = append(testCases, TestCase{
		inputA:   "1010",
		inputB:   "1",
		expected: "1011",
	})
	testCases = append(testCases, TestCase{
		inputA:   "1010",
		inputB:   "1111",
		expected: "11001",
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case no %d", (i+1)), func(tb *testing.T) {
			got := AddBinary(testCase.inputA, testCase.inputB)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case number %d, expected %+v and got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
