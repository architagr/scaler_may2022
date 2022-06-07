package longest_consecutive_one

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestGetLongestConsecutiveOne(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    "01",
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    "111000",
		expected: 3,
	})
	testCases = append(testCases, TestCase{
		input:    "111011101",
		expected: 7,
	})
	testCases = append(testCases, TestCase{
		input:    "1110011101",
		expected: 5,
	})
	testCases = append(testCases, TestCase{
		input:    "1101001100101110",
		expected: 5,
	})
	testCases = append(testCases, TestCase{
		input:    "010100110101",
		expected: 4,
	})
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("testing case no %s", testCase.input), func(tb *testing.T) {
			got := GetLongestConsecutiveOne(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested case number %s, expected %d and got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
