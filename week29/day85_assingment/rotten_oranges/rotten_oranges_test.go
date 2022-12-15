package rotten_oranges

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestRottenOranges(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input:    [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
		expected: 4,
	})
	testCases = append(testCases, TestCase{
		input:    [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
		expected: -1,
	})

	testCases = append(testCases, TestCase{
		input:    [][]int{{0, 2, 1}, {2, 2, 1}, {0, 1, 0}, {2, 1, 1}, {0, 1, 1}, {1, 2, 1}},
		expected: 2,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := RottenOranges(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
