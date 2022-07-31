package flip

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    string
	expected []int
}

func TestCalcFlip(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "010",
		expected: []int{1, 1},
	})
	testCases = append(testCases, TestCase{
		input:    "111",
		expected: []int{},
	})
	testCases = append(testCases, TestCase{
		input:    "000101101",
		expected: []int{1, 3},
	})
	testCases = append(testCases, TestCase{
		input:    "0001001101",
		expected: []int{1, 6},
	})
	testCases = append(testCases, TestCase{
		input:    "0111000100010",
		expected: []int{5, 11},
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := CalcFlip(testCase.input)

			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
