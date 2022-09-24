package permutations_of_a_in_b

import (
	"fmt"
	"testing"
)

func TestPermutationAinB(t *testing.T) {
	type TestCase struct {
		inputA, inputB string
		expected       int
	}
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   "abc",
		inputB:   "abcbacabc",
		expected: 5,
	})
	testCases = append(testCases, TestCase{
		inputA:   "aca",
		inputB:   "acaa",
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		inputA:   "docp",
		inputB:   "aoapeooeoapcpaocecddoocdcqqapeapccc",
		expected: 0,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := PermutationAinB(testCase.inputA, testCase.inputB)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
