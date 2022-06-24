package string_operations

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected string
}

func TestStringOperation(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "AbcaZeoB",
		expected: "bc###bc###",
	})
	testCases = append(testCases, TestCase{
		input:    "AAbcaZeoB",
		expected: "bc###bc###",
	})
	testCases = append(testCases, TestCase{
		input:    "AA",
		expected: "",
	})

	testCases = append(testCases, TestCase{
		input:    "hgUe",
		expected: "hg#hg#",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := StringOperation(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %s but got %s", (i + 1), testCase.expected, got)
			}
		})
	}

}
