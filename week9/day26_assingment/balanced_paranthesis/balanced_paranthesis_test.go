package balanced_paranthesis

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestBalancedParanthesis(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "(())",
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    "()(",
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := BalancedParanthesis(testCase.input)
			if !reflect.DeepEqual(got, testCase.expected) {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
