package dungeon_princess

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    [][]int
	expected int
}

func TestCalculateMinimumHP(t *testing.T) {
	testCases := make([]TestCase, 0)

	testCases = append(testCases, TestCase{
		input: [][]int{
			{-2, -3, 3},
			{-5, -10, 1},
			{10, 30, -5},
		},
		expected: 7,
	})
	testCases = append(testCases, TestCase{
		input: [][]int{
			{1, -1, 0},
			{-1, 1, -11},
			{1, 0, -1},
		},
		expected: 1,
	})

	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing number %d", (i+1)), func(test *testing.T) {
			x := CalculateMinimumHP(testCase.input)

			if x != testCase.expected {
				test.Fatalf("tested %d expected %+v but got %+v", (i + 1), testCase.expected, x)
			}
		})
	}
}
