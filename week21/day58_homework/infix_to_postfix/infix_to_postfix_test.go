package infix_to_postfix

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	expected string
}

func TestInfixToPostfix(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    "x^y/(a*z)+b",
		expected: "xy^az*/b+",
	})
	testCases = append(testCases, TestCase{
		input:    "a+b*(c^d-e)^(f+g*h)-i",
		expected: "abcd^e-fgh*+^*+i-",
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := InfixToPostfix(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
