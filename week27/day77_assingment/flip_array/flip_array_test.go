package flip_array

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []int
	expected int
}

func TestFlipArray(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    []int{15, 10, 6},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		input:    []int{14, 10, 4},
		expected: 1,
	})

	testCases = append(testCases, TestCase{
		input:    []int{8, 4, 5, 7, 6, 2},
		expected: 3,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := FlipArray(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
