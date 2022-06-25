package longest_palindromic_string

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputArr string
	expected string
}

func TestFindLongestPalindrome(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputArr: "abb",
		expected: "bb",
	})
	testCases = append(testCases, TestCase{
		inputArr: "bb",
		expected: "bb",
	})

	testCases = append(testCases, TestCase{
		inputArr: "abc",
		expected: "a",
	})

	testCases = append(testCases, TestCase{
		inputArr: "aaaabaaa",
		expected: "aaabaaa",
	})

	testCases = append(testCases, TestCase{
		inputArr: "caabbaaa",
		expected: "aabbaa",
	})

	testCases = append(testCases, TestCase{
		inputArr: "abcddcbabcd",
		expected: "abcddcba",
	})

	for _, testcase := range testCases {
		t.Run(fmt.Sprintf("testing %v", testcase.inputArr), func(tb *testing.T) {
			got := FindLongestPalindrome(testcase.inputArr)
			if got != testcase.expected {
				tb.Errorf("Tested %+v to  expected %+v but got %+v", testcase.inputArr, testcase.expected, got)
			}
		})
	}
}
