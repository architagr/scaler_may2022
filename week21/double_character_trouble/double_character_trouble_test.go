package double_character_trouble

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected string
}

func TestKReverseLinkedList(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    "ab",
		expected: "ab",
	})
	testCases = append(testCases, TestCase{
		input:    "abccbc",
		expected: "ac",
	})
	testCases = append(testCases, TestCase{
		input:    "abba",
		expected: "",
	})
	testCases = append(testCases, TestCase{
		input:    "aaaaa",
		expected: "a",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := DoubleCharTrouble(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
