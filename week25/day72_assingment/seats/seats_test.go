package seats

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestSeats(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    "..xx.x",
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    "....x..xx...x..",
		expected: 5,
	})
	testCases = append(testCases, TestCase{
		input:    "....xxx",
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		input:    ".x.x.x..x",
		expected: 5,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := seats(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
