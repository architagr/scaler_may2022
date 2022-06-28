package valid_sudoku

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    []string
	expected int
}

func TestValidSudoku(t *testing.T) {

	testCases := make([]TestCase, 0)
	testCases = append(testCases, TestCase{
		input:    []string{"53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"},
		expected: 1,
	})
	testCases = append(testCases, TestCase{
		input:    []string{"53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "5...8..79"},
		expected: 0,
	})
	testCases = append(testCases, TestCase{
		input:    []string{"53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....98.", "...419..5", "....8..79"},
		expected: 0,
	})
	for i, testCase := range testCases {
		t.Run(fmt.Sprintf("testing %d", (i+1)), func(tb *testing.T) {
			got := ValidSudoku(testCase.input)
			if got != testCase.expected {
				tb.Errorf("tested %d expected %d but got %d", (i + 1), testCase.expected, got)
			}
		})
	}
}
