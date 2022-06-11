package bit_compression

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestCompressBits(t *testing.T) {
	testcases := make([]TestCase, 0)
	testcases = append(testcases, TestCase{
		input:    []int{1, 3, 5},
		expected: 7,
	})

	testcases = append(testcases, TestCase{
		input:    []int{1, 1, 1},
		expected: 1,
	})

	for _, testCase := range testcases {
		t.Run(fmt.Sprintf("finding for %d", testCase.input), func(tb *testing.T) {
			got := compressBits(testCase.input)

			if got != testCase.expected {
				tb.Errorf("Tested for %d, expected %d but got %d", testCase.input, testCase.expected, got)
			}
		})
	}
}
