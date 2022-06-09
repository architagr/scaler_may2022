package number_1_bits

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	inputA, expected int
}

func TestFindNumerofOnesBits(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   11,
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case no %d", (i+1)), func(tb *testing.T) {
			got := findNumerofOnesBits(testCase.inputA)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested case number %d, expected %d and got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
