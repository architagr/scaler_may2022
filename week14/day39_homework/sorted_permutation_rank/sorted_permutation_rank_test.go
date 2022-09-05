package sorted_permutation_rank

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestFindRank(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    "acb",
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		input:    "abc",
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		input:    "ecabd",
		expected: 109,
	})

	testCases = append(testCases, TestCase{
		input:    "VIEW",
		expected: 15,
	})

	testCases = append(testCases, TestCase{
		input:    "AaBbCc",
		expected: 51,
	})

	testCases = append(testCases, TestCase{
		input:    "bYRzjSOdiDcGwTUBKerlsXxyqIfHMEAaW",
		expected: 235301,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := FindRank(testCase.input)

			if got != testCase.expected {
				tb.Errorf("testing %d, expected %+v, got %+v", (i + 1), testCase.expected, got)
			}
		})
	}
}
