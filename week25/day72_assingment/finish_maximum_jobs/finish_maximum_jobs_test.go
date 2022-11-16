package finish_maximum_jobs

import (
	"fmt"
	"testing"
)

type TestCase struct {
	inputA, inputB []int
	expected       int
}

func TestMishaAndCandies(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		inputA:   []int{1, 5, 7, 1},
		inputB:   []int{7, 8, 8, 8},
		expected: 2,
	})

	testCases = append(testCases, TestCase{
		inputA:   []int{3, 2, 6},
		inputB:   []int{9, 8, 9},
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := MaxJob(testCase.inputA, testCase.inputB)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
