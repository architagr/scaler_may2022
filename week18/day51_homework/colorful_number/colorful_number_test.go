package colorful_number

import (
	"fmt"
	"reflect"
	"testing"
)

func TestColorfulNumber(t *testing.T) {
	type TestCase struct {
		input    int
		expected int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    236,
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		input:    23,
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ColorfulNumber(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
