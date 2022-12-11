package modified_search

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB []string
	expected       string
}

func TestModifiesSearch(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []string{"data", "circle", "cricket"},
		inputB:   []string{"date", "circel", "crikket", "data", "circl"},
		expected: "10100",
	})

	testCases = append(testCases, TestCase{
		inputA:   []string{"hello", "world"},
		inputB:   []string{"hella", "pello", "pella"},
		expected: "110",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := ModifiesSearch(testCase.inputA, testCase.inputB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
