package letter_phone

import (
	"fmt"
	"reflect"
	"testing"
)

type TestCase struct {
	input    string
	expected []string
}

func TestPermute(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    "012",
		expected: []string{"01a", "01b", "01c"},
	})
	testCases = append(testCases, TestCase{
		input:    "23",
		expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := letterCombinations(testCase.input)

			if !reflect.DeepEqual(x, testCase.expected) {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
