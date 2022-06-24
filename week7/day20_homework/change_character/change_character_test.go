package change_character

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA           string
	expected, inputB int
}

func TestChangeCharacter(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		inputA:   "abcabbccd",
		inputB:   3,
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		inputA:   "abcabbccd",
		inputB:   3,
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ChangeCharacter(testCase.inputA, testCase.inputB)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}

}
