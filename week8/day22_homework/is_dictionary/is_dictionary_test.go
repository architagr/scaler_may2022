package is_dictionary

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA   []string
	inputB   string
	expected int
}

func TestIsDisctionary(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   []string{"hello", "scaler", "interviewbit"},
		inputB:   "adhbcfegskjlponmirqtxwuvzy",
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		inputA:   []string{"fine", "none", "no"},
		inputB:   "qwertyuiopasdfghjklzxcvbnm",
		expected: 0,
	})

	testCases = append(testCases, TestCase{
		inputA:   []string{"ipial", "qjqgt", "vfnue", "vjqfp", "eghva", "ufaeo", "atyqz", "chmxy", "ccvgv", "ghtow"},
		inputB:   "nbpfhmirzqxsjwdoveuacykltg",
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := IsDisctionary(testCase.inputA, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
