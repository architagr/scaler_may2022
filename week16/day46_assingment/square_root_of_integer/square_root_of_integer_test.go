package square_root_of_integer

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, expected int
}

func TestSqrt(t *testing.T) {
	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    4,
		expected: 2,
	})
	testCases = append(testCases, TestCase{
		input:    5,
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := Sqrt(testCase.input)
			if testCase.expected != got {
				tb.Errorf("Tested # %d, where we expected %d but we got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
